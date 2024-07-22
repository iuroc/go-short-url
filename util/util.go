package util

import (
	"encoding/json"
	"errors"
	"github.com/golang-jwt/jwt/v5"
	"go-short-url/database"
	"log"
	"net/http"
	"os"
	"time"
)

// 检查给定的 map 是否缺失必需的键
func CheckMapKeys(inputMap map[string]any, requiredKeys []string) (missing []string) {
	for _, key := range requiredKeys {
		if _, exists := inputMap[key]; !exists {
			missing = append(missing, key)
		}
	}
	return missing
}

// 统一的响应数据结构，可通过自身 Write 方法发送响应
//
// 使用示例：
//
//	util.Response[any]{
//		Success: false,
//		Message: err.Error(),
//	}.Write(w)
//
//	util.Response[any]{
//		Success: true,
//		Message: "注册成功",
//		Data: struct {
//			Username string `json:"username"`
//		}{ Username: username },
//	}.Write(w)
type Response[DataType any] struct {
	// 响应处理是否成功
	Success bool `json:"success"`
	// 响应的主体数据
	Data DataType `json:"data"`
	// 响应的提示内容
	Message string `json:"message"`
}

// 将响应数据构建为 JSON，然后发送到客户端。
//
// 如果 JSON 构建失败，会发送默认的 JSON 字符串。
func (r Response[DataType]) Write(w http.ResponseWriter) {
	body, err := json.Marshal(r)
	if err != nil {
		log.Println(err)
		w.Write([]byte(`{"success":false,"message":"服务器错误","data":null}`))
	} else {
		w.Write(body)
	}
}

// 生成签名后的 Token
func MakeSignedToken(userId int64, username string) string {
	secretKey := []byte(os.Getenv("JWT_KEY"))
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{
		"userId":   userId,
		"username": username,
		"exp":      time.Now().Add(time.Hour * 24).Unix(),
	})
	signedToken, err := token.SignedString(secretKey)
	if err != nil {
		log.Fatalln("[MakeSignedToken]", err)
	}
	return signedToken
}

func CheckSignedToken(tokenString string) (*database.User, error) {
	secretKey := []byte(os.Getenv("JWT_KEY"))
	token, err := jwt.Parse(tokenString, func(token *jwt.Token) (interface{}, error) {
		if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, errors.New("签名方法不正确")
		}
		return secretKey, nil
	})
	if err != nil {
		return nil, err
	}
	if claims, ok := token.Claims.(jwt.MapClaims); ok && token.Valid {
		userId := claims["userId"].(int64)
		userName := claims["user_name"].(string)
		return &database.User{
			Username: userName,
			Id:       userId,
		}, nil
	}
	return nil, errors.New("校验 Token 失败")
}

// 将 UTC ISO 时间转换为本地时间
func ConvertUtcIsoToLocalTime(utcIso string) (string, error) {
	utcTime, err := time.Parse(time.RFC3339, utcIso)
	if err != nil {
		return "", err
	}
	localTime := utcTime.Local()
	return localTime.Format("2006-01-02 15:04:05"), nil
}
