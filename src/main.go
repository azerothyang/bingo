package main

import (
	"conf"
	"github.com/gin-gonic/gin"
	"middleware/holdup"
	"middleware/token"
	"router"
	"runtime"
)

func setupRouter() *gin.Engine {
	// Disable Console Color
	// gin.DisableConsoleColor()
	r := gin.Default()
	r.Use(holdup.CheckHold(), token.HandleToken()) //中间件
	//增加testController 路由
	router.AddUserControllerRoute(r)
	return r
}

func main() {
	//gin.SetMode(gin.ReleaseMode)
	runtime.GOMAXPROCS(runtime.NumCPU())
	r := setupRouter()
	// Listen and Server in 0.0.0.0:8080
	switch {
	case conf.Mode == gin.ReleaseMode:
		r.Run(conf.Addr)
	case conf.Mode == gin.DebugMode:
		r.Run(conf.DevAddr)
	}

}
