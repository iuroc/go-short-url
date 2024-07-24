// 仅用于本项目的工具包
//
// 如果有需要共享给其他项目的功能，请放到共享包中，那里会合理地返回 error 而不是直接结束程序。
package util

import (
	"database/sql"
	"errors"
	"fmt"
	"log"
	"os"
	"regexp"

	_ "github.com/go-sql-driver/mysql"
)

// GetDB 获取数据库连接，并验证 Ping() 情况，如果失败则结束程序。
func GetDB() *sql.DB {
	db, err := sql.Open("mysql", GetDNS())
	if err != nil {
		log.Fatalln("[GetDB] 获取数据库连接失败", err)
	}
	if err = SetTimeZone(db); err != nil {
		log.Fatalln("[GetDB] 初始化设置会话时区失败", err)
	}
	if err = db.Ping(); err != nil {
		log.Fatalln("[GetDB] 获取数据库连接失败", err)
	}
	return db
}

// getDSN 获取数据库连接的 DSN 字符串
func GetDNS() string {
	return fmt.Sprintf("%s:%s@tcp(%s:%s)/%s",
		os.Getenv("DB_USER"),
		os.Getenv("DB_PASSWORD"),
		os.Getenv("DB_HOST"),
		os.Getenv("DB_PORT"),
		os.Getenv("DB_NAME"),
	)
}

// SetTimeZone 设置 MySQL 数据库的时区
//
// 通过读取 TIME_ZONE 环境变量来设置时区，默认是 +08:00。
//
// 将执行下面的 SQL 代码：
//
//	SET time_zone = "+08:00"
func SetTimeZone(db *sql.DB) error {
	timeZone := os.Getenv("TIME_ZONE")
	if timeZone == "" {
		timeZone = "+08:00"
	}
	_, err := db.Exec(fmt.Sprintf(`SET time_zone = "%s"`, timeZone))
	return err
}

// ExecErrorHandler 处理 Exec 执行的异常，示例：
//
//	ExecErrorHandler(stmt.Exec())
func ExecErrorHandler(result sql.Result, err error) (int64, error) {
	if err != nil {
		return -1, err
	}
	if count, err := result.RowsAffected(); err != nil {
		log.Fatalln(err)
	} else if count == 0 {
		return -1, errors.New("受影响的行数为 0")
	}
	return result.LastInsertId()
}

// InitTables 初始化数据表，如果不存在则自动创建。
//
// 包含一次数据库连接，如果数据库连接失败则结束程序。
func InitTables(db *sql.DB) {
	content, err := os.ReadFile("init.sql")
	if err != nil {
		log.Fatalln("[InitTables] 读取 SQL 文件失败", err)
	}
	match := regexp.MustCompile(`(?s)CREATE TABLE .*?(;|$)`)
	result := match.FindAllString(string(content), -1)
	for _, line := range result {
		if _, err = db.Exec(line); err != nil {
			log.Fatalln("[InitTables] 执行 CREATE TABLE 失败", err)
		}
	}
	log.Println("数据表初始化完成")
}
