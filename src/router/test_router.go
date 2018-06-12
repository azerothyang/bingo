package router

import (
	"controller"
	"github.com/gin-gonic/gin"
)

//增加Test控制器的路由
func AddTestControllerRoute(r *gin.Engine) {
	//新建notice
	r.POST("/notice", (&controller.Test{}).Add)

	//获取单个notice详情
	r.GET("/notice/:noticeId", (&controller.Test{}).Get)

	//获取notcie总数
	r.GET("/noticeTotal", (&controller.Test{}).GetTotal)

	r.GET("/redis", (&controller.Test{}).RedisSet)
}
