package router

import (
	"go-short-url/router/handler"
	"net/http"
)

// 用户信息处理路由
func UserRouter() *http.ServeMux {
	router := http.NewServeMux()
	router.HandleFunc("/login", handler.LoginHandler)
	router.HandleFunc("/register", handler.RegisterHandler)
	return router
}
