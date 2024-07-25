package router

import (
	"go-short-url/middleware"
	rulerouter "go-short-url/router/rule"
	userrouter "go-short-url/router/user"
	"net/http"
)

// 路由，匹配全部的路径
func Router() *http.ServeMux {
	router := http.NewServeMux()
	router.HandleFunc("/", redirectHandlerFunc)
	router.Handle("/assets/", http.FileServer(http.Dir("../client/dist")))
	router.Handle("/api/", http.StripPrefix("/api", middleware.ParseFormMiddleware(ApiRouter())))
	return router
}

func ApiRouter() *http.ServeMux {
	router := http.NewServeMux()
	router.Handle("/user/", http.StripPrefix("/user", userrouter.Router()))
	router.Handle("/rule/", http.StripPrefix("/rule", middleware.CheckTokenMiddleware(rulerouter.Router())))
	return router
}
