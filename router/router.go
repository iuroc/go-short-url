package router

import (
	"database/sql"
	"go-short-url/middleware"
	rulerouter "go-short-url/router/rule"
	userrouter "go-short-url/router/user"
	"go-short-url/util"
	"net/http"
	"regexp"
	"strings"
)

// 路由，匹配全部的路径
func Router() *http.ServeMux {
	router := http.NewServeMux()
	router.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		format := r.FormValue("format")
		db := util.GetDB()
		defer db.Close()
		result := regexp.MustCompile(`^/(\w+)(/.*)?`).FindStringSubmatch(r.URL.Path)
		if len(result) == 0 {
			util.Res{Message: "后缀格式不正确"}.Write(w)
			return
		}
		rule, err := rulerouter.SelectRuleBySuffix(db, result[1])
		if err != nil {
			if err == sql.ErrNoRows {
				util.Res{Message: "未找到该短链接规则"}.Write(w)
			} else {
				util.Res{Message: err.Error()}.Write(w)
			}
		} else if format == "json" {
			util.Res{Success: true, Message: "操作成功", Data: rule}.Write(w)
		} else {
			rule.Request = rule.Request + 1
			err := rule.Update(db)
			if err != nil {
				util.Res{Message: err.Error()}.Write(w)
			} else {
				http.Redirect(w, r, strings.TrimRight(rule.Target, "/")+result[2], http.StatusFound)
			}
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
