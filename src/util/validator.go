package util

import (
	"net/url"
	"regexp"
	"strconv"
	"strings"
)

//暂未使用

type Validator struct {
	HasErr  bool              //是否严重有错
	ErrList map[string]string //校验错误列表
	ErrMsg  string            //错误信息
}

type rules map[string]string

//初始化Validator
func NewValidator() *Validator {
	return &Validator{
		HasErr:  false,
		ErrList: make(map[string]string),
		ErrMsg:  "",
	}
}

func (validator *Validator) Validate(form *url.Values, rules rules) *Validator {
	for field, rule := range rules {
		validMethods := strings.Split(rule, "|")
		fieldValidRes := true //单个字段的校验结果
		errorReason := ""     //校验错误原因
		for _, method := range validMethods {
			subMethod := strings.Split(method, ":")
			//TODO 这里只支持一个参数
			arg := ""
			if len(subMethod) == 2 {
				method = subMethod[0]
				arg = subMethod[1]
			}
			if method == "notEmpty" {
				if validator.notEmpty(field, form) == false {
					fieldValidRes = false
					errorReason = method
					break
				}
			}
			if method == "mobile" {
				if validator.mobile(field, form) == false {
					fieldValidRes = false
					errorReason = method
					break
				}
			}
			if method == "password" {
				if validator.password(field, form) == false {
					fieldValidRes = false
					errorReason = method
					break
				}
			}
			if method == "nick" {
				if validator.nick(field, form) == false {
					fieldValidRes = false
					errorReason = method
					break
				}
			}
			if method == "regex" {
				if validator.regex(field, form, arg) == false {
					fieldValidRes = false
					errorReason = method
					break
				}
			}
			if method == "min" {
				if validator.min(field, form, arg) == false {
					fieldValidRes = false
					errorReason = method
					break
				}
			}

			if method == "max" {
				if validator.max(field, form, arg) == false {
					fieldValidRes = false
					errorReason = method
					break
				}
			}
		}
		if fieldValidRes == false {
			validator.HasErr = true
			validator.ErrList[field] = errorReason
			validator.ErrMsg = field + ": " + errorReason + "\r\n"
		}
	}
	return validator
}

//判断参数是否为空
func (*Validator) notEmpty(field string, form *url.Values) bool {
	if len(form.Get(field)) > 0 {
		return true
	}
	return false
}

//手机号验证
func (*Validator) mobile(field string, form *url.Values) bool {
	if ok, _ := regexp.MatchString("^(13[0-9]|14[579]|15[0-3,5-9]|16[6]|17[0135678]|18[0-9]|19[89])\\d{8}$", form.Get(field)); ok {
		return true
	}
	return false
}

//密码验证 密码8-16位数字和字母的组合这两个符号(不能是纯数字或者纯字母)
func (*Validator) password(field string, form *url.Values) bool {
	if ok, _ := regexp.MatchString("^(?![0-9]+$)(?![a-zA-Z]+$)[0-9A-Za-z]{8,16}$", form.Get(field)); ok {
		return true
	}
	return false
}

//用户昵称校验 中文和英文或数字不能有特殊符号长度为2-10位
func (*Validator) nick(field string, form *url.Values) bool {
	if ok, _ := regexp.MatchString("^[a-zA-Z0-9\u4e00-\u9fff]{2,10}$", form.Get(field)); ok {
		return true
	}
	return false
}

//正则校验
func (*Validator) regex(field string, form *url.Values, pattern string) bool {
	if ok, _ := regexp.MatchString(pattern, form.Get(field)); ok {
		return true
	}
	return false
}

//最小只不能小于xx 即验证大于等于arg数值, 下限
func (*Validator) min(field string, form *url.Values, arg string) bool {
	number, err := strconv.Atoi(form.Get(field))
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
func (*Validator) max(field string, form *url.Values, arg string) bool {
	number, err := strconv.Atoi(form.Get(field))
	max, errV := strconv.Atoi(arg)
	if err != nil || errV != nil {
		return false
	}
	if number <= max {
		return true
	}
	return false
}
