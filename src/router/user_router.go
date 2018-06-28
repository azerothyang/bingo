package router

import (
	"controller"
	"github.com/gin-gonic/gin"
)

//增加User控制器的路由
func AddUserControllerRoute(r *gin.Engine) {
	//新建user
	r.POST("/user", (&controller.UserController{}).Add)

	//删除单个user
	r.DELETE("/user/:login_name", (&controller.UserController{}).Del)

	//获取单个user详情
	r.GET("/user/:login_name", (&controller.UserController{}).Get)

	//修改user
	//r.PUT("/user/:login_name", (&controller.Test{}).Get)
}
