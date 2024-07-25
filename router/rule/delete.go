package rule

import (
	"go-short-url/middleware"
	"go-short-url/util"
	"net/http"
	"strconv"
)

func deleteHandlerFunc(w http.ResponseWriter, r *http.Request) {
	id, err := strconv.ParseInt(r.FormValue("id"), 10, 64)
	if err != nil {
		util.Res{Message: "请输入正确格式的 ID"}.Write(w)
		return
	}
	db := util.GetDB()
	defer db.Close()
	token := r.Context().Value(middleware.TokenKey).(*util.TokenInfo)
	err = DeleteById(db, token.UserId, id)
	if err != nil {
		util.Res{Message: "删除失败"}.Write(w)
		return
	}
	util.Res{Message: "删除成功"}.Write(w)
}
