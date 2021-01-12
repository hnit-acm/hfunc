package web

import "github.com/gin-gonic/gin"

type Controller interface {
	RouterRegister(group *gin.RouterGroup)
	RouterGroupName() (name string)
	Middlewares() (middlewares []gin.HandlerFunc)
	Version() string
}
type RouterRegisterFunc func(group *gin.RouterGroup)
type RouterGroupNameFunc func() (name string)
type MiddlewaresFunc func() (middlewares []gin.HandlerFunc)
type VersionFunc func() string

type ControllerFunc func() (RouterRegisterFunc, RouterGroupNameFunc, MiddlewaresFunc, VersionFunc)
