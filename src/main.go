package main

import (
	"common/grace"
	"conf"
	"github.com/gin-gonic/gin"
	"log"
	"middleware/holdup"
	"middleware/token"
	"router"
)

func setupRouter() *gin.Engine {
	// Disable Console Color
	// gin.DisableConsoleColor()
	r := gin.Default()
	r.Use(token.HandleToken(), holdup.CheckHold()) //中间件
	// Ping test
	//r.GET("/ping", func(c *gin.Context) {
	//	c.JSON(http.StatusOK, gin.H{
	//		"code": 200,
	//		"msg": "hello",
	//		"data": "hello world",
	//	})
	//})
	//
	//// Get user value
	//r.GET("/user/:name", func(c *gin.Context) {
	//	user := c.Params.ByName("name")
	//	value, ok := DB[user]
	//	if ok {
	//		c.JSON(200, gin.H{"user": user, "value": value})
	//	} else {
	//		c.JSON(200, gin.H{"user": user, "status": "no value"})
	//	}
	//})
	//
	//Authorized group (uses gin.BasicAuth() middleware)
	// Same than:
	//authorized := r.Group("/")
	//authorized.Use(gin.BasicAuth(gin.Credentials{
	//       "foo":  "bar",
	//       "manu": "123",
	//}))
	//authorized := r.Group("/", gin.BasicAuth(gin.Accounts{
	//	"foo":  "bar", // user:foo password:bar
	//	"manu": "123", // user:manu password:123
	//}))

	//authorized.POST("admin", func(c *gin.Context) {
	//	user := c.MustGet(gin.AuthUserKey).(string)
	//
	//	// Parse JSON
	//	var json struct {
	//		Value string `json:"value" binding:"required"`
	//	}
	//
	//	if c.Bind(&json) == nil {
	//		DB[user] = json.Value
	//		c.JSON(200, gin.H{"status": "ok"})
	//	}
	//})

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
