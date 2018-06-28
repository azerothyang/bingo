package util

import "strconv"

//初始化page和pageSize, 默认page为0, pageSize=10, 返回默认为字符串，方便之后查询拼接字符串
func InitPageAndPageSize(page string, orgPageSize string) (offset string, size string) {
	p, err := strconv.ParseInt(page, 10, 11)
	ps, errz := strconv.ParseInt(orgPageSize, 10, 11)
	if err != nil || errz != nil {
		return "0", "10"
	}
	if p < 0 || ps < 0 {
		return "0", "10"
	}
	off := int(p * ps)
	sz := int(ps)
	return strconv.Itoa(off), strconv.Itoa(sz)
}
