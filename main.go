package main

import (
	"github.com/joho/godotenv"
	"go-short-url/database"
	"go-short-url/router"
	"log"
	"net/http"
)

func main() {
	err := godotenv.Load()
	if err != nil {
		log.Fatalln(err)
	}
	// 获取数据库连接
	db := database.GetDB()
	defer db.Close()
	// 初始化数据表
	database.InitTables(db, "init.sql")
	// 初始化管理员账号
	database.InitAdminUser(db)
	log.Println("服务启动成功 👉 http://127.0.0.1:9090")
	http.ListenAndServe("127.0.0.1:9090", router.MainRouter())
}
