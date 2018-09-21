package test

import (
	"fmt"
	"testing"
	"util/validate"
)

func TestValidator(t *testing.T) {
	form := map[string]string{
		"code":  "",
		"msg":   "",
		"phone": "1832085.5",
	}

	rules := map[string]string{
		"code":  "",
		"msg":   "",
		"phone": "numeric",
	}
	validator := validate.New()
	validator.Validate(&form, rules)
	fmt.Println(validator.HasErr, validator.ErrList)
}
