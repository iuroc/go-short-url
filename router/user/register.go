package user

import (
	"go-short-url/util"
	"golang.org/x/crypto/bcrypt"
	"log"
	"net/http"
	"strings"
)

func registerHandlerFunc(w http.ResponseWriter, r *http.Request) {
	username := strings.TrimSpace(r.FormValue("username"))
	password := strings.TrimSpace(r.FormValue("password"))
	if err := util.CheckUsernameAndPasswordFormat(username, password); err != nil {
		util.Res{Message: err.Error()}.Write(w)
		return
	}
	db := util.GetDB()
	defer db.Close()
	hash, err := bcrypt.GenerateFromPassword([]byte(password), bcrypt.DefaultCost)
	if err != nil {
		log.Fatalln("[registerHandlerFunc]", err)
	}
	user, err := User{
		Username:     username,
		PasswordHash: string(hash),
		Role:         "user",
	}.Insert(db)
	if err != nil {
		if strings.Contains(err.Error(), "Duplicate entry") {
			util.Res{Message: "用户名已存在，请换一个重试"}.Write(w)
		} else {
			util.Res{Message: err.Error()}.Write(w)
		}
	} else {
		util.Res{Success: true, Message: "注册成功", Data: user}.Write(w)
	}
}
