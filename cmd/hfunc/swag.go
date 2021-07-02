package main

import (
	"encoding/json"
	"errors"
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/hnit-acm/hfunc/hapi"
	"github.com/hnit-acm/hfunc/hserver/hhttp"
	"github.com/lucas-clemente/quic-go/http3"
	swaggerFiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"
	"github.com/swaggo/swag"
	"io/ioutil"
	"net/http"
	"net/http/httputil"
	"net/url"
	"os"
	"regexp"
	"strings"
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
	return hapi.ServeAny(hhttp.WithHandler(r), hhttp.WithAddr(":"+port))
}
