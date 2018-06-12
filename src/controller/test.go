package controller

import (
	"common/validation"
	"github.com/gin-gonic/gin"
	"model"
	"net/http"
	"util"
	"util/status"
)

type Test struct {
	BaseController
}

//新建notice
func (b *Test) Add(c *gin.Context) {
	var notice model.Notice
	//请求参数进行数据绑定
	if err := c.ShouldBind(&notice); err != nil {
		c.JSON(200, gin.H{
			"code": status.VALIDATE_FAIL,
			"msg":  status.GetReason(status.VALIDATE_FAIL),
		})
		return
	}
	//数据校验
	valid := validation.Validation{}
	valid.Valid(&notice)
	if valid.HasErrors() {
		// 如果有错误信息，证明验证没通过
		errorList := make(map[string]string)
		for _, err := range valid.Errors {
			errorList[err.Key] = err.Message
		}
		c.JSON(200, gin.H{
			"code": status.VALIDATE_FAIL,
			"msg":  status.GetReason(status.VALIDATE_FAIL),
			"data": errorList,
		})
		return
	}
	noticeCode := util.RandomStr(32) //生成noticeCode
	//验证通过, 将数据写入数据库
	data := map[string]string{
		"notice_content": notice.NoticeContent,
		"notice_type":    notice.NoticeType,
		"notice_code":    noticeCode,
	}
	lastId, err := (&model.NoticeModel{}).Add(&data)
	if err != nil {
		c.JSON(http.StatusInternalServerError, "")
	}
	c.JSON(http.StatusOK, gin.H{
		"code": http.StatusOK,
		"msg":  nil,
		"data": lastId,
	})
}

func (b *Test) Get(c *gin.Context) {
	noticeId := c.Param("noticeId")
	notice, err := (&model.NoticeModel{}).GetById(noticeId, "notice_code, notice_content, create_time")
	if err != nil {

	}
	c.JSON(http.StatusOK, gin.H{
		"code": http.StatusOK,
		"msg":  nil,
		"data": *notice,
	})
}

//获取数据库总数的例子
func (b *Test) GetTotal(c *gin.Context) {
	res := (&model.NoticeModel{}).GetAllCount()
	c.JSON(http.StatusOK, gin.H{
		"code": http.StatusOK,
		"msg":  nil,
		"data": res,
	})
}

func (b *Test) RedisSet(c *gin.Context) {
	str, _ := Redis.Gett("x")
	c.JSON(http.StatusOK, gin.H{
		"code": http.StatusOK,
		"msg":  nil,
		"data": str,
	})
}
