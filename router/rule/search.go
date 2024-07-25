package rule

import (
	"go-short-url/middleware"
	"go-short-url/util"
	"net/http"
	"strconv"
	"strings"
)

func searchHandlerFunc(w http.ResponseWriter, r *http.Request) {
	pageStr := strings.TrimSpace(r.FormValue("page"))
	pageSizeStr := strings.TrimSpace(r.FormValue("pageSize"))
	keyword := strings.TrimSpace(r.FormValue("keyword"))
	if pageStr == "" {
		pageStr = "0"
	}
	if pageSizeStr == "" {
		pageSizeStr = "72"
	}
	page, err := strconv.Atoi(pageStr)
	pageSize, err2 := strconv.Atoi(pageSizeStr)
	if err != nil || err2 != nil {
		util.Res{Message: "请输入正确格式的 page 和 pageSize"}.Write(w)
		return
	}
	db := util.GetDB()
	defer db.Close()
	token := r.Context().Value(middleware.TokenKey).(*util.TokenInfo)
	rules, err := SearchRules(db, token.UserId, keyword, page, pageSize)
	if err != nil {
		util.Res{Message: err.Error()}.Write(w)
	} else {
		util.Res{Success: true, Message: "查询成功", Data: rules}.Write(w)
	}
}
