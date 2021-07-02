package hapi

import (
	"github.com/hnit-acm/hfunc/hserver/hhttp"
	"github.com/lucas-clemente/quic-go/http3"
	"net/http"
)

//ServeAny support http1.1 http2 http3
func ServeAny(ops ...hhttp.Option) error {
	httpServer := hhttp.New(ops...)
	quicServer := &http3.Server{
		Server: httpServer,
	}
	handler := httpServer.Handler
	httpServer.Handler = http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		quicServer.SetQuicHeaders(w.Header())
		handler.ServeHTTP(w, r)
	})

	hErr := make(chan error)
	qErr := make(chan error)
	go func() {
		hErr <- hhttp.ListenTLS(httpServer)
	}()
	go func() {
		qErr <- quicServer.ListenAndServe()
	}()

	select {
	case err := <-hErr:
		quicServer.Close()
		return err
	case err := <-qErr:
		return err
	}
}
