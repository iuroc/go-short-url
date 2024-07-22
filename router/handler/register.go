package handler

import (
	"go-short-url/database"
	"go-short-url/util"
	"log"
	"net/http"
	"strings"
)

func RegisterHandler(w http.ResponseWriter, r *http.Request) {
	if r.ParseForm() != nil {
		return
	}
	username := strings.TrimSpace(r.FormValue("username"))
	password := strings.TrimSpace(r.FormValue("password"))
	var err error
	// 检查用户名格式是否正确
	if err = util.CheckUsernameFormat(username); err != nil {
		util.Response[any]{
			Success: false,
			Message: err.Error(),
		}.Write(w)
		return
	}
	// 检查密码格式是否正确
	if err = util.CheckPasswordFormat(password); err != nil {
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
	}.Insert(db)); err != nil {
		errorMessage := "注册失败"
		log.Println("[RegisterHandler] 注册失败", err)
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
