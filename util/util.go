package util

import (
	"encoding/json"
	"errors"
	"golang.org/x/crypto/bcrypt"
	"log"
	"net/http"
	"regexp"
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

// 检查用户名格式是否正确
func CheckUsernameFormat(username string) error {
	if regexp.MustCompile(`^\w{3,20}$`).MatchString(username) {
		return nil
	} else {
		return errors.New("用户名格式错误，要求 3-20 位，可使用数字、字母、下划线")
	}
}

// 检查密码格式是否正确
func CheckPasswordFormat(password string) error {
	if regexp.MustCompile(`^[\x00-\x7F]{8,20}$`).MatchString(password) {
		return nil
	} else {
		return errors.New("密码格式错误，要求 8-20 位，可使用数字、字母、特殊符号")
	}
}

// 将明文密码转换为 bcrypt 哈希密码
func HashPassword(password string) (string, error) {
	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(password), bcrypt.DefaultCost)
	if err != nil {
		return "", nil
	}
	return string(hashedPassword), nil
}

// 验证明文密码和哈希密码是否匹配
func CheckPasswordHash(password string, hash string) bool {
	return bcrypt.CompareHashAndPassword([]byte(hash), []byte(password)) == nil
}
