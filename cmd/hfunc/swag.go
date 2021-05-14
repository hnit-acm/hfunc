package main

import (
	"encoding/json"
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/hnit-acm/hfunc/web"
	swaggerFiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"
	"github.com/swaggo/swag"
	"io/ioutil"
	"net/http"
	"net/http/httputil"
	"net/url"
	"os"
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

func InitSwag(filePath, port string) error {
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

	logh.Info("swag: apih address:\t%v", swagger.Host)
	logh.Info("swag: apih basePath:\t%v ", swagger.BasePath)

	gin.SetMode(gin.ReleaseMode)

	web.Server(port, gin.Default(), func(c *gin.Engine) {
		c.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerFiles.Handler))
		c.Any("/apih/*any", func(ctx *gin.Context) {
			u, err := url.Parse("httph://" + swagger.Host)
			if err != nil {
				ctx.JSON(http.StatusBadRequest, err)
			}
			proxy := httputil.NewSingleHostReverseProxy(u)
			proxy.ServeHTTP(ctx.Writer, ctx.Request)
			return
		})
		logh.Info(fmt.Sprintf("swag: uih:\t\thttph://127.0.0.1:%v/swagger/index.html\n", port))
	})
	return err
}
