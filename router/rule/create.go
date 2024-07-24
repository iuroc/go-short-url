package rule

import (
	"go-short-url/middleware"
	"go-short-url/util"
	"net/http"
	"net/url"
	"strings"
	"time"
)

func createHandlerFunc(w http.ResponseWriter, r *http.Request) {
	suffix := strings.TrimSpace(r.FormValue("suffix"))
	target := strings.TrimSpace(r.FormValue("target"))

	token := r.Context().Value(middleware.TokenKey).(*util.TokenInfo)
	if suffix == "" || target == "" {
		util.Res{Message: "后缀和目标网址不能为空"}.Write(w)
		return
	}
	if len(suffix) > 20 {
		util.Res{Message: "后缀长度不能超过 20 个字符"}.Write(w)
		return
	}
	_url, err := url.Parse(target)
	if err != nil || (_url.Scheme != "http" && _url.Scheme != "https") {
		util.Res{Message: "请输入正确格式的目标网址"}.Write(w)
		return
	}
	if _url.Hostname() == r.URL.Hostname() {
		util.Res{Message: "不支持缩短当前网址"}.Write(w)
		return
	}

	expiresString := strings.TrimSpace(r.FormValue("expires"))
	var expires *time.Time
	if expiresString != "" {
		t, err := time.Parse(time.RFC3339, expiresString)
		if err != nil {
			util.Res{Message: "expires 字段必须是 RFC3339 格式的日期时间"}.Write(w)
			return
		}
		t = t.In(util.GetLocation())
		expires = &t
	}
	rule := Rule{
		Suffix:  suffix,
		Target:  target,
		UserId:  token.UserID,
		Expires: expires,
	}
	db := util.GetDB()
	defer db.Close()
	insertId, err := rule.Insert(db)
	if err != nil {
		if strings.Contains(err.Error(), "Duplicate entry") {
			util.Res{Message: "后缀已被使用，请换一个重试"}.Write(w)
		} else {
			util.Res{Message: err.Error()}.Write(w)
		}
		return
	}
	util.Res{Success: true, Message: "创建成功", Data: insertId}.Write(w)
}
