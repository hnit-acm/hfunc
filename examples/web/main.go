package main

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/hnit-acm/hfunc/apih"
)

type testController struct {
}

func (testController) HelloFunc() apih.HandleFunc {
	return func() (httpMethod, routeUri, version string, handlerFunc gin.HandlerFunc) {
		return "GET", "helloFunc", "", func(c *gin.Context) {
			fmt.Println("runing hello world")
			c.JSON(200, "hello world")
			return
		}
	}
}

func (testController) Hello(c *gin.Context) {
	fmt.Println("runing hello world")
	c.JSON(200, "hello world")
	return
}

func Middleware(c *gin.Context) {
	fmt.Println("before")
	c.Next()
	fmt.Println("after")
}

func (t testController) RouterRegister(group *gin.RouterGroup) {
	group.GET("/hello", t.Hello)
}

func (t testController) RouterGroupName() (name string) {
	return "/test"
}

func (t testController) Version() string {
	return "v1"
}

func (t testController) Middlewares() (middlewares []gin.HandlerFunc) {
	return []gin.HandlerFunc{
		Middleware,
	}
}

var test2Controller = apih.ControllerFunc(
	func() (apih.RouterRegisterFunc, apih.RouterGroupNameFunc, apih.MiddlewaresFunc, apih.VersionFunc) {
		return func(group *gin.RouterGroup) {
				group.GET("hello", func(context *gin.Context) {
					context.JSON(200, "heihei")
				})
			}, func() (name string) {
				return "test2"
			}, func() (middlewares []gin.HandlerFunc) {
				return []gin.HandlerFunc{
					Middleware,
				}
			}, func() string {
				return "v2"
			}
	},
)

func main() {
	apih.Server("8080", nil, func(c *gin.Engine) {
		apih.Register(c,
			func(engine *gin.Engine) *gin.RouterGroup {
				return engine.Group("/apih")
			},
			testController{},
		)

		apih.RegisterFunc(
			c,
			func(engine *gin.Engine) *gin.RouterGroup {
				return engine.Group("/apih")
			},
			test2Controller,
		)

		apih.RegisterHandleFunc(
			c,
			func(engine *gin.Engine) *gin.RouterGroup {
				return engine.Group("/apih")
			},
			testController{},
		)
	})
}
