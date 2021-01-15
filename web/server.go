package web

import (
	"context"
	"log"
	"net/http"
	"os"
	"os/signal"
	"syscall"
	"time"

	"github.com/gin-gonic/gin"
	serverhttp "github.com/hnit-acm/hfunc/server/http"
)

// Server 启动gin server
// port 端口
// regFunc 注册的路由函数
func Server(port string, g *gin.Engine, regFunc func(c *gin.Engine)) {
	if g == nil {
		g = gin.Default()
	}
	if regFunc != nil {
		regFunc(g)
	}

	httpServer := serverhttp.NewServer("tcp", ":"+port, serverhttp.Handler(g))

	startCtx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	go func() {
		if err := httpServer.Start(startCtx); err != nil && err != http.ErrServerClosed {
			log.Fatalf("listen: %s\n", err)
		}
	}()

	quit := make(chan os.Signal)
	signal.Notify(quit, syscall.SIGINT, syscall.SIGTERM)
	<-quit

	log.Println("Shutting down server...")

	stopCtx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()
	httpServer.Stop(stopCtx)
	log.Println("Server exiting")
}
