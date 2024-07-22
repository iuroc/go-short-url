package router

import (
	"go-short-url/router/rule"
	"go-short-url/router/user"
	"net/http"
)

func MainRouter() *http.ServeMux {
	router := http.NewServeMux()
	router.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		w.Write([]byte(r.URL.Path))
	})
	router.Handle("/api/user/", http.StripPrefix("/api/user", userrouter.UserRouter()))
	router.Handle("/api/rule/", http.StripPrefix("/api/rule", rulerouter.URLRouter()))
	return router
}
