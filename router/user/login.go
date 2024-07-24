package user

import (
	"go-short-url/util"
	"net/http"
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
		user := SelectUserByNameAndPassword(db, username, password)
		if user == nil {
			util.Res{Message: "查询失败"}.Write(w)
		} else {
			util.Res{Data: user, Message: "查询成功"}.Write(w)
		}
	} else {
		// 验证 Cookie，当字段未找到时，会返回错误
		token, err := r.Cookie("token")
		if err != nil {
			// Cookie 中 Token 字段缺失
			util.Res{Message: "token 校验失败"}.Write(w)
			return
		}
		util.Res{Data: token.Value}.Write(w)
	}
}
