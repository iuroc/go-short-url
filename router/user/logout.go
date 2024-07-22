package userrouter

import (
	"go-short-url/util"
	"net/http"
	"time"
)

func LogoutHandler(w http.ResponseWriter, r *http.Request) {
	http.SetCookie(w, &http.Cookie{
		Name:    "token",
		Value:   "",
		Expires: time.Now(),
	})
	util.Response[any]{
		Success: true,
		Message: "退出登录成功，已清除 Cookie",
	}.Write(w)
}
