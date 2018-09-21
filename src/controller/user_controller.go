package controller

import (
	"github.com/gin-gonic/gin"
	"net/http"
)

type UserController struct {
	BaseController
}

//获取用户列表
func (u *UserController) GetList(c *gin.Context) {
	users := userModel.GetList()
	c.JSON(http.StatusOK, gin.H{
		"code": http.StatusOK,
		"msg":  nil,
		"data": users,
	})
}
