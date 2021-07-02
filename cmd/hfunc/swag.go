package main

import (
	"crypto/rand"
	"crypto/rsa"
	"crypto/tls"
	"crypto/x509"
	"crypto/x509/pkix"
	"encoding/json"
	"encoding/pem"
	"errors"
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/lucas-clemente/quic-go/http3"
	swaggerFiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"
	"github.com/swaggo/swag"
	"io/ioutil"
	"math/big"
	"net"
	"net/http"
	"net/http/httputil"
	"net/url"
	"os"
	"regexp"
	"strings"
	"time"
)

type s struct {
	filePath string
	err      error
}

func (s s) ReadDoc() string {
	bytes, err := os.ReadFile(s.filePath)
	if err != nil {
		s.err = err
		return ""
	}
	var m map[string]interface{}
	err = json.Unmarshal(bytes, &m)
	if err != nil {
		s.err = err
		return ""
	}
	m["host"] = ""
	bytes, err = json.Marshal(m)
	return string(bytes)
}

type SwagInfo struct {
	Info struct {
		Version     string `json:"version,omitempty"`
		Title       string `json:"title,omitempty"`
		Description string `json:"description,omitempty"`
	} `json:"info"`
	Swagger  string `json:"swagger"`
	Host     string `json:"host"`
	BasePath string `json:"basePath"`
}

func InitSwag(filePath, port, rewrite string) error {
	logh.Info("swag: filePath:\t\t%v", filePath)
	logh.Info("swag: port:\t\t%v", port)

	bytes, err := ioutil.ReadFile(filePath)
	if err != nil {
		return err
	}

	swagger := SwagInfo{}
	err = json.Unmarshal(bytes, &swagger)
	if err != nil {
		return err
	}
	swag.Register(swag.Name, &s{
		filePath: filePath,
	})

	logh.Info("swag: api address:\t%v", swagger.Host)
	logh.Info("swag: api basePath:\t%v ", swagger.BasePath)

	rewriteCmd := strings.Split(rewrite, " ")
	if len(rewriteCmd) != 2 && rewrite != "" {
		return errors.New("swag: url rewrite format error")
	}
	logh.Info("swag: url rewrite:\t%v", rewrite)

	gin.SetMode(gin.ReleaseMode)
	r := gin.Default()
	r.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerFiles.Handler))
	r.Any("/api/*any", func(ctx *gin.Context) {
		u, err := url.Parse("http://" + swagger.Host)
		if err != nil {
			ctx.JSON(http.StatusBadRequest, err)
			return
		}
		if len(rewriteCmd) == 2 {
			reg, err := regexp.Compile(rewriteCmd[0])
			if err != nil {
				ctx.JSON(http.StatusBadRequest, err)
				return
			}
			ctx.Request.URL.Path = reg.ReplaceAllString(ctx.Request.URL.Path, rewriteCmd[1])
		}
		proxy := httputil.NewSingleHostReverseProxy(u)
		proxy.Transport = &http3.RoundTripper{}
		proxy.ServeHTTP(ctx.Writer, ctx.Request)
		return
	})
	logh.Info(fmt.Sprintf("swag: ui:\t\thttps://127.0.0.1:%v/swagger/index.html\n", port))

	return server(":"+port, r, generateTLSConfig())
}

func generateTLSConfig() *tls.Config {
	max := new(big.Int).Lsh(big.NewInt(1), 128)
	serialNumber, _ := rand.Int(rand.Reader, max)
	subject := pkix.Name{
		Country:            []string{"CN"},
		Province:           []string{"BeiJing"},
		Organization:       []string{"Devops"},
		OrganizationalUnit: []string{"certDevops"},
		CommonName:         "127.0.0.1",
	}

	template := x509.Certificate{
		SerialNumber: serialNumber,
		Subject:      subject,
		NotBefore:    time.Now(),
		NotAfter:     time.Now().Add(365 * 24 * time.Hour),
		KeyUsage:     x509.KeyUsageKeyEncipherment | x509.KeyUsageDigitalSignature,
		ExtKeyUsage:  []x509.ExtKeyUsage{x509.ExtKeyUsageServerAuth},
		IPAddresses:  []net.IP{net.ParseIP("127.0.0.1")},
	}

	pk, _ := rsa.GenerateKey(rand.Reader, 2048)

	derBytes, _ := x509.CreateCertificate(rand.Reader, &template, &template, &pk.PublicKey, pk)
	certPEM := pem.EncodeToMemory(&pem.Block{Type: "CERTIFICATE", Bytes: derBytes})

	keyPEM := pem.EncodeToMemory(&pem.Block{Type: "RSA PRIVATE KEY", Bytes: x509.MarshalPKCS1PrivateKey(pk)})

	tlsCert, err := tls.X509KeyPair(certPEM, keyPEM)
	if err != nil {
		panic(err)
	}
	return &tls.Config{
		Certificates: []tls.Certificate{tlsCert},
		NextProtos:   []string{"hfunc"},
	}
}

func server(addr string, handler http.Handler, config *tls.Config) error {
	var err error
	udpAddr, err := net.ResolveUDPAddr("udp", addr)
	if err != nil {
		return err
	}
	udpConn, err := net.ListenUDP("udp", udpAddr)
	if err != nil {
		return err
	}
	defer udpConn.Close()

	tcpAddr, err := net.ResolveTCPAddr("tcp", addr)
	if err != nil {
		return err
	}
	tcpConn, err := net.ListenTCP("tcp", tcpAddr)
	if err != nil {
		return err
	}
	defer tcpConn.Close()

	tlsConn := tls.NewListener(tcpConn, config)
	defer tlsConn.Close()

	httpServer := &http.Server{
		Addr:      addr,
		TLSConfig: config,
	}

	quicServer := &http3.Server{
		Server: httpServer,
	}

	if handler == nil {
		handler = http.DefaultServeMux
	}
	httpServer.Handler = http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		quicServer.SetQuicHeaders(w.Header())
		handler.ServeHTTP(w, r)
	})

	hErr := make(chan error)
	qErr := make(chan error)
	go func() {
		hErr <- httpServer.Serve(tlsConn)
	}()
	go func() {
		qErr <- quicServer.Serve(udpConn)
	}()

	select {
	case err := <-hErr:
		quicServer.Close()
		return err
	case err := <-qErr:
		return err
	}
}
