package userrouter

import (
	"net/http"
)

// 用户信息处理路由
func UserRouter() *http.ServeMux {
	router := http.NewServeMux()
	router.HandleFunc("/login", LoginHandler)
	router.HandleFunc("/register", RegisterHandler)
	router.HandleFunc("/logout", LogoutHandler)
	return router
}
