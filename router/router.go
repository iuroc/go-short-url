package router

import (
	"go-short-url/middleware"
	rulerouter "go-short-url/router/rule"
	userrouter "go-short-url/router/user"
	"go-short-url/util"
	"net/http"
)

// 路由，匹配全部的路径
func Router() *http.ServeMux {
	router := http.NewServeMux()
	router.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		db := util.GetDB()
		defer db.Close()
		rule, err := rulerouter.SelectTargetBySuffix(db, "iuroc")
		if err != nil {
			util.Res{Message: err.Error()}.Write(w)
		} else {
			util.Res{Success: true, Message: "操作成功", Data: rule}.Write(w)
		}
	})
	router.Handle("/api/", http.StripPrefix("/api", middleware.ParseFormMiddleware(ApiRouter())))
	return router
}

func ApiRouter() *http.ServeMux {
	router := http.NewServeMux()
	router.Handle("/user/", http.StripPrefix("/user", userrouter.Router()))
	router.Handle("/rule/", http.StripPrefix("/rule", middleware.CheckTokenMiddleware(rulerouter.Router())))
	return router
}
