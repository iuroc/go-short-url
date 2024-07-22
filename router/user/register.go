package userrouter

import (
	"go-short-url/database"
	"go-short-url/util"
	"net/http"
	"strings"
)

func RegisterHandler(w http.ResponseWriter, r *http.Request) {
	if err := r.ParseForm(); err != nil {
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
	// 注册账号开始
	passwordHash := util.HashPassword(password)
	db := database.GetDB()
	defer db.Close()
	if user, err := (database.User{
		Username: username,
		Password: passwordHash,
		Role:     "user",
	}.Insert(db)); err != nil {
		errorMessage := "注册失败"
		if strings.Contains(err.Error(), "Duplicate entry") {
			errorMessage = "用户名已存在，请换一个再注册"
		}
		util.Response[any]{
			Success: false,
			Message: errorMessage,
		}.Write(w)
		return
	} else {
		util.Response[*database.User]{
			Success: true,
			Message: "注册成功",
			Data:    user,
		}.Write(w)
	}
}
