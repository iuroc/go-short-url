package rulerouter

import (
	"net/http"
)

// 链接信息处理路由
func URLRouter() *http.ServeMux {
	router := http.NewServeMux()

	router.HandleFunc("/create", CreateHandler)

	return router
}
