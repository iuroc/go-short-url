package user

import (
	"go-short-url/util"
	"net/http"
)

func loginHandler(w http.ResponseWriter, r *http.Request) {
	// if query.Has("username") && query.Has("password") {
	// 	username := query.Get("username")
	// 	password := query.Get("password")
	// 	if err := util.CheckUsernameAndPasswordFormat(username, password); err != nil {
	// 		util.Res{}.Write(w)
	// 		return
	// 	}
	// }
	util.Res{}.Write(w)
}
