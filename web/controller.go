package web

import (
	"github.com/gin-gonic/gin"
	"reflect"
)

type Controller interface {
	RouterRegister(group *gin.RouterGroup)
	RouterGroupName() (name string)
	Middlewares() (middlewares []gin.HandlerFunc)
	Version() string
}

type HandleFunc func() (httpMethod, routeUri, version string, handlerFunc gin.HandlerFunc)
type NewHandleFunc func() HandleFunc

var _emptyHandleFunc = HandleFunc(
	func() (httpMethod, routeUri, version string, handlerFunc gin.HandlerFunc) {
		return "", "", "", nil
	},
)
var _emptyNewHandleFunc = NewHandleFunc(
	func() HandleFunc {
		return func() (httpMethod, routeUri, version string, handlerFunc gin.HandlerFunc) {
			return "", "", "", nil
		}
	},
)

type RouterRegisterFunc func(group *gin.RouterGroup)
type RouterGroupNameFunc func() (name string)
type MiddlewaresFunc func() (middlewares []gin.HandlerFunc)
type VersionFunc func() string

type ControllerFunc func() (RouterRegisterFunc, RouterGroupNameFunc, MiddlewaresFunc, VersionFunc)

// RegisterHandleFunc
// 该方法可以自动注入controller中按规范定义的handelfunc HandleFunc NewHandleFunc
// 注意： handlefunc的版本优先级比controller高
func RegisterHandleFunc(router *gin.Engine, routeReg func(*gin.Engine) *gin.RouterGroup, cs ...Controller) {
	if routeReg != nil {
		r := routeReg(router)
		if r != nil {
			for _, c := range cs {
				v := reflect.ValueOf(c)
				for i := 0; i < v.NumMethod(); i++ {
					//e:=v.Method(i).Func
					switch {
					case v.Method(i).Type().ConvertibleTo(reflect.TypeOf(_emptyNewHandleFunc)):
						{
							outs := v.Method(i).Call(nil)
							f := outs[0].Interface().(HandleFunc)
							httpMethod, routeUri, version, handlerFunc := f()
							if version != "" {
								r.Group(version).Group(c.RouterGroupName(), c.Middlewares()...).Handle(httpMethod, routeUri, handlerFunc)
							}
							r.Group(c.Version()).Group(c.RouterGroupName(), c.Middlewares()...).Handle(httpMethod, routeUri, handlerFunc)
						}
					case v.Method(i).Type().ConvertibleTo(reflect.TypeOf(_emptyHandleFunc)):
						{
							outs := v.Method(i).Call(nil)
							httpMethod, routeUri, version, handlerFunc := outs[0].Interface().(string),
								outs[1].Interface().(string),
								outs[2].Interface().(string),
								outs[3].Interface().(gin.HandlerFunc)
							if version != "" {
								r.Group(version).Group(c.RouterGroupName(), c.Middlewares()...).Handle(httpMethod, routeUri, handlerFunc)
							}
							r.Group(c.Version()).Group(c.RouterGroupName(), c.Middlewares()...).Handle(httpMethod, routeUri, handlerFunc)
						}
					}
				}
			}
			return
		}
	}
	for _, c := range cs {
		v := reflect.ValueOf(c)
		for i := 0; i < v.NumMethod(); i++ {
			//e:=v.Method(i).Func
			switch {
			case v.Method(i).Type().ConvertibleTo(reflect.TypeOf(_emptyNewHandleFunc)):
				{
					outs := v.Method(i).Call(nil)
					f := outs[0].Interface().(HandleFunc)
					httpMethod, routeUri, version, handlerFunc := f()
					if version != "" {
						router.Group(version).Group(c.RouterGroupName(), c.Middlewares()...).Handle(httpMethod, routeUri, handlerFunc)
					}
					router.Group(c.Version()).Group(c.RouterGroupName(), c.Middlewares()...).Handle(httpMethod, routeUri, handlerFunc)
				}
			case v.Method(i).Type().ConvertibleTo(reflect.TypeOf(_emptyHandleFunc)):
				{
					outs := v.Method(i).Call(nil)
					httpMethod, routeUri, version, handlerFunc := outs[0].Interface().(string),
						outs[1].Interface().(string),
						outs[2].Interface().(string),
						outs[3].Interface().(gin.HandlerFunc)
					if version != "" {
						router.Group(version).Group(c.RouterGroupName(), c.Middlewares()...).Handle(httpMethod, routeUri, handlerFunc)
					}
					router.Group(c.Version()).Group(c.RouterGroupName(), c.Middlewares()...).Handle(httpMethod, routeUri, handlerFunc)
				}
			}
		}
	}

}
