package status

const (
	VALIDATE_FAIL = 1001 //数据校验失败
	INTER_ERROR   = 500  //数据校验失败d
)

var Reason = map[int]string{
	VALIDATE_FAIL: "数据格式错误",
	INTER_ERROR:   "内部错误",
}

//获取错误码错误原因
func GetReason(errorCode int) string {
	value, ok := Reason[errorCode]
	if ok {
		return value
	}
	return ""
}
