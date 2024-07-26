package user

import (
	"go-short-url/util"
	"golang.org/x/crypto/bcrypt"
	"net/http"
	"time"
)

func loginHandlerFunc(w http.ResponseWriter, r *http.Request) {
	if r.Form.Has("username") || r.Form.Has("password") {
		// 验证表单
		username := r.Form.Get("username")
		password := r.Form.Get("password")
		if username == "" {
			util.Res{Message: "用户名不能为空"}.Write(w)
			return
		}
		if len(username) > 50 {
			util.Res{Message: "用户名过长"}.Write(w)
			return
		}
		if password == "" {
			util.Res{Message: "密码不能为空"}.Write(w)
			return
		}
		if len(password) > 50 {
			util.Res{Message: "密码过长"}.Write(w)
			return
		}
		db := util.GetDB()
		defer db.Close()
		user := SelectUserByName(db, username)
		if user == nil || bcrypt.CompareHashAndPassword([]byte(user.PasswordHash), []byte(password)) != nil {
			util.Res{Message: "用户名或密码错误"}.Write(w)
		} else {
			cookie := http.Cookie{
				Name:    "token",
				Value:   util.MakeToken(user.Id, user.Username),
				Expires: time.Now().Add(time.Hour * 24 * 60),
			}
			http.SetCookie(w, &cookie)
			util.Res{Success: true, Data: user, Message: "登录成功"}.Write(w)
		}
	} else {
		// 验证 Cookie，当字段未找到时，会返回错误
		tokenCookie, err := r.Cookie("token")
		if err != nil {
			util.Res{Message: "token 校验失败"}.Write(w)
			return
		}
		tokenInfo, err := util.CheckToken(tokenCookie.Value)
		if err != nil {
			// Cookie 中 Token 字段缺失
			util.Res{Message: "token 校验失败"}.Write(w)
		} else {
			db := util.GetDB()
			defer db.Close()
			user := SelectUserByName(db, tokenInfo.Username)
			util.Res{Success: true, Message: "token 校验成功", Data: user}.Write(w)
		}
	}
}
