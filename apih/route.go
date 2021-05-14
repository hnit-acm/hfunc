package apih

import "github.com/gin-gonic/gin"

func Register(router *gin.Engine, routeReg func(*gin.Engine) *gin.RouterGroup, cs ...Controller) {
	if routeReg != nil {
		r := routeReg(router)
		if r != nil {
			for _, c := range cs {
				group := r.Group(c.Version()).Group(c.RouterGroupName(), c.Middlewares()...)
				c.RouterRegister(group)
			}
			return
		}
	}
	for _, c := range cs {
		group := router.Group(c.Version()).Group(c.RouterGroupName(), c.Middlewares()...)
		c.RouterRegister(group)
	}
}

func RegisterFunc(router *gin.Engine, routeReg func(*gin.Engine) *gin.RouterGroup, cs ...ControllerFunc) {
	if routeReg != nil {
		r := routeReg(router)
		if r != nil {
			for _, c := range cs {
				RouterRegister, RouterGroupName, Middlewares, Version := c()
				group := r.Group(Version()).Group(RouterGroupName(), Middlewares()...)
				RouterRegister(group)
			}
			return
		}
	}
	for _, c := range cs {
		RouterRegister, RouterGroupName, Middlewares, Version := c()
		group := router.Group(Version()).Group(RouterGroupName(), Middlewares()...)
		RouterRegister(group)
	}
}
