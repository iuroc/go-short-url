package user

import (
	"go-short-url/util"
	"net/http"
	"time"

	"golang.org/x/crypto/bcrypt"
)

func loginHandlerFunc(w http.ResponseWriter, r *http.Request) {
	if r.Form.Has("username") || r.Form.Has("password") {
		// 验证表单
		username := r.Form.Get("username")
		password := r.Form.Get("password")
		if err := util.CheckUsernameAndPasswordFormat(username, password); err != nil {
			util.Res{Message: err.Error()}.Write(w)
			return
		}
		db := util.GetDB()
		defer db.Close()
		user := SelectUserByName(db, username)
		if user == nil || bcrypt.CompareHashAndPassword([]byte(user.PasswordHash), []byte(password)) != nil {
			util.Res{Message: "用户名或密码错误"}.Write(w)
		} else {
			cookie := http.Cookie{
				Name:  "token",
				Value: util.MakeToken(user.Id, user.Username),
				Expires: time.Now().Add(time.Hour * 24 * 60),
			}
			http.SetCookie(w, &cookie)
			util.Res{Success: true, Data: user, Message: "登录成功"}.Write(w)
		}
	} else {
		// 验证 Cookie，当字段未找到时，会返回错误
		token, err := r.Cookie("token")
		if err != nil || !util.CheckToken(token.Value) {
			// Cookie 中 Token 字段缺失
			util.Res{Message: "token 校验失败"}.Write(w)
		} else {
			util.Res{Success: true, Message: "token 校验成功"}.Write(w)
		}
	}
}
