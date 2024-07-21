package router

import "net/http"

// 链接信息处理路由
func URLRouter() *http.ServeMux {
	router := http.NewServeMux()
	return router
}
