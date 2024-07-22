package handler

import (
	"go-short-url/database"
	"go-short-url/util"
	"net/http"
	"strings"
	"time"
)

// 处理用户登录
//
// 1. 账号密码登录
//
// 2. JWT 登录
func LoginHandler(w http.ResponseWriter, r *http.Request) {
	username := strings.TrimSpace(r.FormValue("username"))
	password := strings.TrimSpace(r.FormValue("password"))

	if username != "" && password != "" {
		// 账号密码验证
		if err := r.ParseForm(); err != nil {
			// 表单数据解析失败
			util.Response[any]{
				Success: false,
				Message: err.Error(),
			}.Write(w)
			return
		}

		username := strings.TrimSpace(r.FormValue("username"))
		password := strings.TrimSpace(r.FormValue("password"))
		if err := util.CheckUsernameAndPasswordFormat(username, password); err != nil {
			util.Response[any]{
				Success: false,
				Message: err.Error(),
			}.Write(w)
			return
		}

		db := database.GetDB()
		defer db.Close()
		if !database.CheckLogin(db, username, password) {
			util.Response[string]{
				Success: false,
				Message: "登录失败，用户名或密码错误",
			}.Write(w)
			return
		}

		token := util.MakeSignedToken(username)

		http.SetCookie(w, &http.Cookie{
			Name:     "token",
			Value:    token,
			Path:     "/",
			Expires:  time.Now().Add(24 * time.Hour),
			HttpOnly: true,
		})

		w.Header().Set("token", token)
		util.Response[string]{
			Success: true,
			Message: "登录成功",
		}.Write(w)
	} else {
		token := ""
		for _, cookie := range r.Cookies() {
			if cookie.Name == "token" {
				token = cookie.Value
				break
			}
		}
		// 使用 JWT 验证
		if util.CheckSignedToken(token) {
			util.Response[string]{
				Success: true,
				Message: "token 校验成功",
			}.Write(w)
		} else {
			util.Response[string]{
				Success: false,
				Message: "token 校验失败",
			}.Write(w)
		}
	}
}
