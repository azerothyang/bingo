package main

import (
	"common/grace"
	"conf"
	"github.com/gin-gonic/gin"
	"log"
	"middleware/ckcache"
	"middleware/holdup"
	"middleware/token"
	"router"
)

func setupRouter() *gin.Engine {
	// Disable Console Color
	// gin.DisableConsoleColor()
	r := gin.Default()
	r.Use(token.HandleToken(), holdup.CheckHold(), ckcache.CheckCache()) //中间件
	//增加testController 路由
	router.AddTestControllerRoute(r)

	return r
}

func main() {
	//gin.SetMode(gin.ReleaseMode)

	route := setupRouter()
	// Listen and Server in 0.0.0.0:8080
	//r.Run(":8080")
	setUpSrv(route)
}

//启动服务器, 优化支持平滑重启
func setUpSrv(router *gin.Engine) {
	var srv *grace.Server
	if conf.Mode == gin.DebugMode {
		srv = grace.NewServer(
			conf.DevAddr,
			router,
		)
	} else {
		srv = grace.NewServer(
			conf.Addr,
			router,
		)
	}
	if err := srv.ListenAndServe(); err != nil {
		log.Fatalln("listen: " + err.Error())
	}

}
