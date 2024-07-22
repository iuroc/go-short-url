package rulerouter

import (
	"go-short-url/database"
	"go-short-url/util"
	"net/http"
	"strings"
)

func CreateHandler(w http.ResponseWriter, r *http.Request) {
	err := r.ParseForm()
	if err != nil {
		util.Response[any]{Message: err.Error(), Success: false}.Write(w)
		return
	}
	rule, err := GetValues(r)
	db := database.GetDB()
	defer db.Close()
}

func GetValues(r *http.Request) (*database.Rule, error) {
	suffix := strings.TrimSpace(r.FormValue("suffix"))
	target := strings.TrimSpace(r.FormValue("target"))
	userId := strings.TrimSpace(r.FormValue("userId"))
	expires := strings.TrimSpace(r.FormValue("expires"))
	return nil, nil
}
