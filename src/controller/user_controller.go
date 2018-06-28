package controller

import (
	"github.com/gin-gonic/gin"
	"model"
	"net/http"
	"common/bingo"
	"util/validate"
	"util"
	"util/status"
)

type UserController struct {
	BaseController
}

//新建用户
func (u *UserController) Add(c *gin.Context) {
	request, _ := bingo.MergeRequest(c.Request, c.Params)
	validator := validate.New()
	rules := map[string]string{
		"login_name": "notEmpty|maxLength:64",
		"pwd": "notEmpty|maxLength:64",
	}
	validator.Validate(request, rules)
	//如果验证没通过直接退出
	if validator.HasErr {
		c.JSON(http.StatusOK, gin.H{
			"code": http.StatusBadRequest,
			"msg":  http.StatusText(http.StatusBadRequest),
			"data": validator.ErrList,
		})
		return
	}
	util.OnlyCols([]string{"login_name", "pwd"}, request)
	//判断login_name是否重复, 如果重复直接返回
	user, err := (&model.UserModel{}).GetByLoginName((*request)["login_name"], "id")
	//如果查询到了对应的用户, 则用户名重复了
	if user.Id != 0 || err != nil {
		c.JSON(http.StatusOK, gin.H{
			"code": http.StatusBadRequest,
			"msg":  status.GetReason(status.LOGIN_NAME_REPEAT),
			"data": nil,
		})
		return
	}
	//验证通过, 将数据写入数据库
	lastId, err := (&model.UserModel{}).Add(request)
	if err != nil {
		c.JSON(http.StatusOK, gin.H{
			"code": http.StatusInternalServerError,
			"msg":  http.StatusText(http.StatusInternalServerError),
			"data": nil,
		})
		return
	}
	c.JSON(http.StatusOK, gin.H{
		"code": http.StatusOK,
		"msg":  nil,
		"data": lastId,
	})
}

//获取单个用户信息
func (u *UserController) Get(c *gin.Context) {
	request, _ := bingo.MergeRequest(c.Request, c.Params)
	validator := validate.New()
	rules := map[string]string{
		"login_name": "notEmpty|maxLength:64",
	}
	validator.Validate(request, rules)
	//如果验证没通过直接退出
	if validator.HasErr {
		c.JSON(http.StatusOK, gin.H{
			"code": http.StatusBadRequest,
			"msg":  http.StatusText(http.StatusBadRequest),
			"data": validator.ErrList,
		})
		return
	}
	user, err := (&model.UserModel{}).GetByLoginName((*request)["login_name"], "id, pwd")
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"code": http.StatusBadRequest,
			"msg":  http.StatusText(http.StatusBadRequest),
			"data": nil,
		})
		return
	}
	c.JSON(http.StatusOK, gin.H{
		"code": http.StatusOK,
		"msg":  nil,
		"data": *user,
	})
}

//获取单个用户信息
func (u *UserController) Del(c *gin.Context) {
	request, _ := bingo.MergeRequest(c.Request, c.Params)
	validator := validate.New()
	rules := map[string]string{
		"login_name": "notEmpty|maxLength:64",
	}
	validator.Validate(request, rules)
	//如果验证没通过直接退出
	if validator.HasErr {
		c.JSON(http.StatusOK, gin.H{
			"code": http.StatusBadRequest,
			"msg":  http.StatusText(http.StatusBadRequest),
			"data": validator.ErrList,
		})
		return
	}
	res, err := (&model.UserModel{}).DeleteUserByLoginName((*request)["login_name"])
	if err != nil {
		c.JSON(http.StatusOK, gin.H{
			"code": http.StatusInternalServerError,
			"msg": nil,
			"data": nil,
		})
		return
	}
	c.JSON(http.StatusOK, gin.H{
		"code": http.StatusOK,
		"msg": nil,
		"data": res,
	})
}
