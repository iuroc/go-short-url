package database

import (
	"database/sql"
	"errors"
	"fmt"
	"go-short-url/util"
	"log"
	"os"
	"strings"

	"github.com/golang-jwt/jwt/v5"
)

// 获取数据库连接，并验证 Ping() 情况。
func GetDB() *sql.DB {
	db, err := sql.Open("mysql", fmt.Sprintf("%s:%s@tcp(%s:%s)/%s",
		os.Getenv("DB_USER"),
		os.Getenv("DB_PASSWORD"),
		os.Getenv("DB_HOST"),
		os.Getenv("DB_PORT"),
		os.Getenv("DB_NAME"),
	))
	if err != nil {
		log.Fatalln("[GetDB] 获取数据库连接失败", err)
	}
	if _, err = db.Exec("SET time_zone = '+08:00'"); err != nil {
		log.Fatalln("[GetDB] 初始化设置会话时区失败", err)
	}
	if err = db.Ping(); err != nil {
		log.Fatalln("[GetDB] 获取数据库连接失败", err)
	}
	return db
}

// 初始化数据表，不存在数据表则自动创建。
func InitTables(db *sql.DB, sqlFilePath string) {
	bytes, err := os.ReadFile(sqlFilePath)
	if err != nil {
		log.Fatalln("[InitTables]", err)
	}
	// 校验管理员密码格式
	if err = util.CheckPasswordFormat(os.Getenv("ROOT_PASSWORD")); err != nil {
		log.Fatalln("[InitTables-6] 管理员初始密码格式错误", err.Error())
	}
	querys := strings.Split(string(bytes), ";")
	// 批量初始化创建表
	for _, query := range querys {
		if strings.TrimSpace(query) != "" {
			if _, err = db.Exec(query); err != nil {
				log.Fatalln("[InitTables]", err)
			}
		}
	}
	log.Println("初始化数据表成功")
}

// 初始化管理员账户
func InitAdminUser(db *sql.DB) {
	rootUsername := strings.TrimSpace(os.Getenv("ROOT_USERNAME"))
	rootPassword := strings.TrimSpace(os.Getenv("ROOT_PASSWORD"))
	err := util.CheckUsernameFormat(rootUsername)
	if err != nil {
		log.Fatalln("[InitAdminUser-1]", err)
	}
	err = util.CheckPasswordFormat(rootPassword)
	if err != nil {
		log.Fatalln("[InitAdminUser-2]", err)
	}
	passwordHash := util.HashPassword(rootPassword)
	// 创建管理员账户
	if exists, _ := AdminExists(db); !exists {
		if _, err = (User{
			Username: rootUsername,
			Password: passwordHash,
			Role:     "admin",
		}.Insert(db)); err != nil {
			log.Fatalln("[InitAdminUser-4]", err)
		}
		log.Println("管理员账号自动创建成功")
	}
}

func CheckSignedToken(tokenString string) (*User, error) {
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
		return &User{
			Username: userName,
			Id:       userId,
		}, nil
	}
	return nil, errors.New("校验 Token 失败")
}