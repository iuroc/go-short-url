package main

import (
	"go-short-url/router"
	"go-short-url/util"
	"log"
	"net/http"

	"github.com/joho/godotenv"
)

func main() {
	// 检查环境变量字段完整性
	CheckEnvs()
	// 初始化数据表
	util.InitTables()
	// 启动服务器
	log.Println("服务启动成功 👉 http://127.0.0.1:9091")
	http.ListenAndServe("127.0.0.1:9091", router.Router())
}

// CheckEnvs 检查环境变量的完整性，不完整则结束程序。
func CheckEnvs() {
	if err := godotenv.Load(); err != nil {
		log.Fatalln(err)
	}
	util.RequireEnvs("DB_PASSWORD", "DB_NAME", "JWT_KEY", "ROOT_USERNAME", "ROOT_PASSWORD")
}
