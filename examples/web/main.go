package main

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/hnit-acm/hfunc/hapi"
	"github.com/hnit-acm/hfunc/hserver/hhttp"
)

type testController struct {
}

func (testController) HelloFunc() hapi.HandleFunc {
	return func() (httpMethod, routeUri, version string, handlerFunc gin.HandlerFunc) {
		return "GET", "helloFuncerr", "", func(c *gin.Context) {
			fmt.Println("runing hello world")
			hapi.JsonResponseErr(c, AuthErr)
			return
		}
	}
}

func (testController) HelloFuncok() hapi.HandleFunc {
	return func() (httpMethod, routeUri, version string, handlerFunc gin.HandlerFunc) {
		return "GET", "helloFuncok", "", func(c *gin.Context) {
			fmt.Println("runing hello world")
			hapi.JsonResponseOk(c, "ok", hapi.WithCode(200))
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

var test2Controller = hapi.ControllerFunc(
	func() (hapi.RouterRegisterFunc, hapi.RouterGroupNameFunc, hapi.MiddlewaresFunc, hapi.VersionFunc) {
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

var (
	AuthErr = hapi.NewCodeErr(200, "")
)

func main() {
	c := gin.Default()
	hapi.Register(c,
		func(engine *gin.Engine) *gin.RouterGroup {
			return engine.Group("/hapi")
		},
		testController{},
	)

	hapi.RegisterFunc(
		c,
		func(engine *gin.Engine) *gin.RouterGroup {
			return engine.Group("/hapi")
		},
		test2Controller,
	)

	hapi.RegisterHandleFunc(
		c,
		func(engine *gin.Engine) *gin.RouterGroup {
			return engine.Group("/hapi")
		},
		testController{},
	)
	hapi.ServeAny(hhttp.WithHandler(c))
}
