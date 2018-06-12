package util

import "strconv"

//初始化page和pageSize, 默认page为0, pageSize=10, 返回默认为字符串，方便之后查询拼接字符串
func InitPageAndPageSize(page *int, orgPageSize *int) (offset string, size string) {
	if *page < 0 {
		*page = 0
	}
	if *orgPageSize <= 0 {
		*orgPageSize = 10
	}
	offset = strconv.Itoa((*page) * (*orgPageSize))
	return offset, strconv.Itoa(*orgPageSize)
}
