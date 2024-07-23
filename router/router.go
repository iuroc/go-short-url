package router

import (
	rulerouter "go-short-url/router/rule"
	userrouter "go-short-url/router/user"
	"go-short-url/util"
	"net/http"
)

// 路由，匹配全部的路径
func Router() *http.ServeMux {
	router := http.NewServeMux()
	router.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		w.Write([]byte("此时即将重定向"))
	})
	router.Handle("/api/", http.StripPrefix("/api", ParseFormMiddleware(ApiRouter())))
	return router
}

func ApiRouter() *http.ServeMux {
	router := http.NewServeMux()
	router.Handle("/user/", http.StripPrefix("/user", userrouter.Router()))
	router.Handle("/rule/", http.StripPrefix("/rule", rulerouter.Router()))
	return router
}

func ParseFormMiddleware(handler http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		if r.ParseForm() != nil {
			util.Res{Message: "参数错误"}.Write(w)
		} else {
			handler.ServeHTTP(w, r)
		}
	})
}
