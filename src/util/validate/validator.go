package validate

import (
	"regexp"
	"strconv"
	"strings"
	"unicode/utf8"
)

type Validator struct {
	HasErr  bool              //是否严重有错
	ErrList map[string]string //校验错误列表
}

type rules map[string]string

//初始化Validator
func New() *Validator {
	return &Validator{
		HasErr:  false,
		ErrList: make(map[string]string),
	}
}

/**
form直接传bingo.MergeRequest中处理返回的键值对请求数据
rules参数形式:, map[string]string{
	"param1": "numeric|min:10",
	"param1": "numeric|min:10",
}
*/
func (validator *Validator) Validate(form *map[string]string, rules rules) *Validator {
	for field, rule := range rules {
		valis := strings.Split(rule, "|")
		validator.HasErr = false //单个字段校验结果是否有错
		for _, vali := range valis {
			if vali == "" {
				continue
			}
			subValis := strings.Split(vali, ":")
			//这里只支持一个参数
			arg := ""
			method := vali
			//如果子校验规则带有参数则再做处理
			if len(subValis) == 2 {
				method = subValis[0]
				arg = subValis[1]
			}
			if method == "notEmpty" {
				if validator.notEmpty(field, form) == false {
					validator.HasErr = true
					validator.ErrList[field] = "不能为空"
					continue
				}
			}
			if method == "mobile" {
				if validator.mobile(field, form) == false {
					validator.HasErr = true
					validator.ErrList[field] = "格式错误"
					continue
				}
			}
			if method == "password" {
				if validator.password(field, form) == false {
					validator.HasErr = true
					validator.ErrList[field] = "格式错误"
					continue
				}
			}
			if method == "nick" {
				if validator.nick(field, form) == false {
					validator.HasErr = true
					validator.ErrList[field] = "格式错误"
					continue
				}
			}
			if method == "regex" {
				if validator.regex(field, form, arg) == false {
					validator.HasErr = true
					validator.ErrList[field] = "格式错误"
					continue
				}
			}
			if method == "min" {
				if validator.min(field, form, arg) == false {
					validator.HasErr = true
					validator.ErrList[field] = "值小于规定值"
					continue
				}
			}

			// max:3
			if method == "max" {
				if validator.max(field, form, arg) == false {
					validator.HasErr = true
					validator.ErrList[field] = "大于指定值"
					continue
				}
			}

			// minLength
			if method == "minLength" {
				if validator.minLength(field, form, arg) == false {
					validator.HasErr = true
					validator.ErrList[field] = "小于最小长度"
					continue
				}
			}

			// maxLength
			if method == "maxLength" {
				if validator.maxLength(field, form, arg) == false {
					validator.HasErr = true
					validator.ErrList[field] = "超过最大长度"
					continue
				}
			}

			// numeric
			if method == "numeric" {
				if validator.numeric(field, form) == false {
					validator.HasErr = true
					validator.ErrList[field] = "非数值型"
					continue
				}
			}
		}
	}
	return validator
}

//判断参数是否为空
func (*Validator) notEmpty(field string, form *map[string]string) bool {
	v, ok := (*form)[field]
	if !ok || v == "" {
		return false
	}
	return true
}

//手机号验证
func (*Validator) mobile(field string, form *map[string]string) bool {
	v, _ := (*form)[field]
	if ok, _ := regexp.MatchString("^(13[0-9]|14[579]|15[0-3,5-9]|16[6]|17[0135678]|18[0-9]|19[89])\\d{8}$", v); ok {
		return true
	}
	return false
}

//密码验证 密码8-16位数字和字母的组合这两个符号(不能是纯数字或者纯字母)
func (*Validator) password(field string, form *map[string]string) bool {
	v, _ := (*form)[field]
	if ok, _ := regexp.MatchString("^(?![0-9]+$)(?![a-zA-Z]+$)[0-9A-Za-z]{8,16}$", v); ok {
		return true
	}
	return false
}

//用户昵称校验 中文和英文或数字不能有特殊符号长度为2-10位
func (*Validator) nick(field string, form *map[string]string) bool {
	v, _ := (*form)[field]
	if ok, _ := regexp.MatchString("^[a-zA-Z0-9\u4e00-\u9fff]{2,10}$", v); ok {
		return true
	}
	return false
}

//正则校验
func (*Validator) regex(field string, form *map[string]string, pattern string) bool {
	v, _ := (*form)[field]
	if ok, _ := regexp.MatchString(pattern, v); ok {
		return true
	}
	return false
}

//最小只不能小于xx 即验证大于等于arg数值, 下限
func (*Validator) min(field string, form *map[string]string, arg string) bool {
	v, _ := (*form)[field]
	number, err := strconv.Atoi(v)
	min, errV := strconv.Atoi(arg)
	if err != nil || errV != nil {
		return false
	}
	if number >= min {
		return true
	}
	return false
}

//最大值不能大于xx 即验证小于等于 arg数值, 上限
func (*Validator) max(field string, form *map[string]string, arg string) bool {
	v, _ := (*form)[field]
	number, err := strconv.Atoi(v)
	max, errV := strconv.Atoi(arg)
	if err != nil || errV != nil {
		return false
	}
	if number <= max {
		return true
	}
	return false
}

//数值型
func (*Validator) numeric(field string, form *map[string]string) bool {
	v, _ := (*form)[field]
	_, err := strconv.ParseFloat(v, 64)
	if err != nil {
		return false
	}
	return true
}

//字符串最大长度
func (*Validator) maxLength(field string, form *map[string]string, arg string) bool {
	v, _ := (*form)[field]
	maxLength, err := strconv.Atoi(arg)
	if err != nil {
		return false
	}
	if utf8.RuneCountInString(v) <= maxLength {
		return true
	}
	return false
}

//字符串最小长度
func (*Validator) minLength(field string, form *map[string]string, arg string) bool {
	v, _ := (*form)[field]
	minLength, err := strconv.Atoi(arg)
	if err != nil {
		return false
	}
	if utf8.RuneCountInString(v) >= minLength {
		return true
	}
	return false
}
