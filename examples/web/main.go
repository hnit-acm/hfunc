package main

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/hnit-acm/hfunc/web"
)

type testController struct {
}

func (testController) HelloFunc() web.HandleFunc {
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

var test2Controller = web.ControllerFunc(
	func() (web.RouterRegisterFunc, web.RouterGroupNameFunc, web.MiddlewaresFunc, web.VersionFunc) {
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
	web.Server("8080", nil, func(c *gin.Engine) {
		web.Register(c,
			func(engine *gin.Engine) *gin.RouterGroup {
				return engine.Group("/api")
			},
			testController{},
		)

		web.RegisterFunc(
			c,
			func(engine *gin.Engine) *gin.RouterGroup {
				return engine.Group("/api")
			},
			test2Controller,
		)

		web.RegisterHandleFunc(
			c,
			func(engine *gin.Engine) *gin.RouterGroup {
				return engine.Group("/api")
			},
			testController{},
		)
	})
}
