package util

import (
	"errors"
	"regexp"
)

// 检查密码格式是否正确
func CheckPasswordFormat(password string) error {
	if regexp.MustCompile(`^[\x00-\x7F]{8,20}$`).MatchString(password) {
		return nil
	} else {
		return errors.New("密码格式错误，要求 8-20 位，可使用数字、字母、特殊符号")
	}
}

// 检查用户名格式是否正确
func CheckUsernameFormat(username string) error {
	if regexp.MustCompile(`^\w{3,20}$`).MatchString(username) {
		return nil
	} else {
		return errors.New("用户名格式错误，要求 3-20 位，可使用数字、字母、下划线")
	}
}

// 校验用户名和密码的格式
func CheckUsernameAndPasswordFormat(username string, password string) error {
	var err error
	// 检查用户名格式是否正确
	if err = CheckUsernameFormat(username); err != nil {
		return err
	}
	// 检查密码格式是否正确
	if err = CheckPasswordFormat(password); err != nil {
		return err
	}
	return nil
}