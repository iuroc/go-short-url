package mixin

import (
	"database/sql"
	"errors"
	"log"
	"regexp"

	"golang.org/x/crypto/bcrypt"
)

// 处理执行 Exec 后的异常，返回 LastInsertId 和需要抛出的错误
func HandleExecError(result sql.Result, err error) (int64, error) {
	if err != nil {
		return 0, errors.New("操作失败，" + err.Error())
	}
	if rows, err := result.RowsAffected(); err != nil {
		log.Fatalln("[HandleExecError-1]", err)
	} else if rows == 0 {
		return 0, errors.New("操作失败，受影响的行数为 0")
	}
	id, err := result.LastInsertId()
	if err != nil {
		log.Fatalln("[HandleExecError-2]", err)
	}
	return id, nil
}


// 验证明文密码和哈希密码是否匹配
func CheckPasswordHash(password string, hash string) bool {
	return bcrypt.CompareHashAndPassword([]byte(hash), []byte(password)) == nil
}


// 检查密码格式是否正确
func CheckPasswordFormat(password string) error {
	if regexp.MustCompile(`^[\x00-\x7F]{8,20}$`).MatchString(password) {
		return nil
	} else {
		return errors.New("密码格式错误，要求 8-20 位，可使用数字、字母、特殊符号")
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


// 检查用户名格式是否正确
func CheckUsernameFormat(username string) error {
	if regexp.MustCompile(`^\w{3,20}$`).MatchString(username) {
		return nil
	} else {
		return errors.New("用户名格式错误，要求 3-20 位，可使用数字、字母、下划线")
	}
}



// 将明文密码转换为 bcrypt 哈希密码
func HashPassword(password string) string {
	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(password), bcrypt.DefaultCost)
	if err != nil {
		log.Fatalln("[HashPassword]", err)
	}
	return string(hashedPassword)
}

