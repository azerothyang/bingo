package bingo

import (
	"net/http"
	"github.com/gin-gonic/gin"
)

//处理前端请求中携带的所有参数, 合并到map中
func MergeRequest(req *http.Request, params gin.Params) (*map[string]string, error){
	mergeReq := make(map[string]string)
	err := req.ParseForm()
	if err != nil {
		return &mergeReq, err
	}
	for k := range req.PostForm{
		mergeReq[k] = req.PostForm.Get(k)
	}
	for _, param := range params {
		mergeReq[param.Key] = param.Value
	}
	return &mergeReq, err
}
