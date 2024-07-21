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
	db, err := database.GetDB()
	if err != nil {
		log.Fatalln("[main] 获取数据库连接失败", err)
	} else {
		log.Println("获取数据库连接成功")
	}
	defer db.Close()
	// 初始化数据表
	err = database.InitTables(db, "init.sql")
	if err != nil {
		log.Fatalln("[main] 初始化数据表失败", err)
	} else {
		log.Println("初始化数据表成功")
	}
	log.Println("服务启动成功 👉 http://127.0.0.1:9090")
	http.ListenAndServe("127.0.0.1:9090", router.MainRouter())
}
