package database

import (
	"database/sql"
	"fmt"
	"go-short-url/util"
	"log"
	"os"
	"strings"

	_ "github.com/go-sql-driver/mysql"
)

// 获取数据库连接，并验证 Ping() 情况。
func GetDB() (*sql.DB, error) {
	db, err := sql.Open("mysql", fmt.Sprintf("%s:%s@tcp(%s:%s)/%s",
		os.Getenv("DB_USER"),
		os.Getenv("DB_PASSWORD"),
		os.Getenv("DB_HOST"),
		os.Getenv("DB_PORT"),
		os.Getenv("DB_NAME"),
	))
	if err != nil {
		return db, err
	}
	return db, db.Ping()
}

// 初始化数据表，不存在数据表则自动创建。
func InitTables(db *sql.DB, sqlFilePath string) error {
	bytes, err := os.ReadFile(sqlFilePath)
	if err != nil {
		return err
	}
	// 校验管理员密码格式
	if err = util.CheckPasswordFormat(os.Getenv("ROOT_PASSWORD")); err != nil {
		log.Fatalln("[InitTables] 管理员初始密码格式错误", err.Error())
	}
	querys := strings.Split(string(bytes), ";")
	for _, query := range querys {
		if strings.TrimSpace(query) != "" {
			if _, err = db.Exec(query); err != nil {
				return err
			}
		}
	}
	return err
}
