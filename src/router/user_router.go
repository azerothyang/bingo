package router

import (
	"controller"
	"github.com/gin-gonic/gin"
)

//增加User控制器的路由
func AddUserControllerRoute(r *gin.Engine) {
	//获取user列表
	r.GET("/users", (&controller.UserController{}).GetList)
}
