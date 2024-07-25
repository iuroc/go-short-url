package user

import (
	"go-short-url/util"
	"net/http"
	"time"
)

func logoutHandlerFunc(w http.ResponseWriter, r *http.Request) {
	http.SetCookie(w, &http.Cookie{
		Name:    "token",
		Value:   "",
		Expires: time.Now(),
	})
	util.Res{Message: "退出登录成功", Success: true}.Write(w)
}
