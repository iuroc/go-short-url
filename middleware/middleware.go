package middleware

import (
	"context"
	"go-short-url/util"
	"net/http"
)

// ParseFormMiddleware 中间件，用于解析 application/x-www-form-urlencoded 格式的 Body
//
// 在 Body 格式错误时，会调用：
//
//	util.Res{Message: "参数错误"}.Write(w)
func ParseFormMiddleware(handler http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		if r.ParseForm() != nil {
			util.Res{Message: "参数错误"}.Write(w)
		} else {
			handler.ServeHTTP(w, r)
		}
	})
}

type Key string

const (
	TokenKey Key = "token"
)

func CheckTokenMiddleware(handler http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		// 验证 Cookie，当字段未找到时，会返回错误
		tokenCookie, err := r.Cookie("token")
		if err != nil {
			util.Res{Message: "token 校验失败"}.Write(w)
			return
		}
		token, err := util.CheckToken(tokenCookie.Value)
		if err != nil {
			// Cookie 中 Token 字段缺失
			util.Res{Message: "token 校验失败"}.Write(w)
		} else {
			r = r.WithContext(context.WithValue(r.Context(), TokenKey, token))
			handler.ServeHTTP(w, r)
		}
	})
}
