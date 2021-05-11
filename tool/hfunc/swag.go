package main

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/hnit-acm/hfunc/web"
	swaggerFiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"
	"github.com/swaggo/swag"
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
	fmt.Println(string(bytes))
	return string(bytes)
}

func InitSwag(filePath, port string) error {
	_, err := os.Open(filePath)
	if err != nil {
		return err
	}
	swag.Register(swag.Name, &s{
		filePath: filePath,
	})
	gin.SetMode(gin.ReleaseMode)
	web.Server(port, gin.New(), func(c *gin.Engine) {
		c.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerFiles.Handler))
	})
	return nil
}
