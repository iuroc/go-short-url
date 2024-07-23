package router

import (
	"go-short-url/router/rule"
	"go-short-url/router/user"
	"net/http"
)

// 路由，匹配全部的路径
func Router() *http.ServeMux {
	router := http.NewServeMux()
	router.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		w.Write([]byte("此时即将重定向"))
	})
	router.Handle("/api/user/", http.StripPrefix("/api/user", userrouter.Router()))
	router.Handle("/api/rule/", http.StripPrefix("/api/rule", rulerouter.Router()))
	return router
}
