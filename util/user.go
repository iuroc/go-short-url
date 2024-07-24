package util

import (
	"errors"
	"log"
	"os"
	"regexp"
	"time"

	"github.com/golang-jwt/jwt/v5"
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

// MakeToken 生成 JWT 字符串，从环境变量读取 JWT_KEY。
func MakeToken(id int64, username string) string {
	key := []byte(os.Getenv("JWT_KEY"))
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{
		"userId":   id,
		"username": username,
		"exp":      time.Now().Add(time.Hour * 24 * 60).Unix(),
	})
	tokenString, err := token.SignedString(key)
	if err != nil {
		log.Println(err)
	}
	return tokenString
}

type TokenInfo struct {
	Token    *jwt.Token
	UserID   int64
	Username string
}

func CheckToken(tokenString string) (*TokenInfo, error) {
	key := []byte(os.Getenv("JWT_KEY"))
	token, err := jwt.Parse(tokenString, func(token *jwt.Token) (interface{}, error) {
		if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, errors.New("未知的签名算法")
		}
		return key, nil
	})
	if err != nil {
		return nil, err
	}
	_, ok := token.Claims.(jwt.MapClaims)
	var userId int64
	var username string
	if chaims, ok := token.Claims.(jwt.MapClaims); ok {
		userId = int64(chaims["userId"].(float64))
		username = chaims["username"].(string)
	}
	if ok && token.Valid {
		return &TokenInfo{
			Token:    token,
			UserID:   userId,
			Username: username,
		}, nil
	} else {
		return nil, errors.New("校验失败")
	}
}
