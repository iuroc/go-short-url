package rule

import (
	"go-short-url/middleware"
	"go-short-url/util"
	"net/http"
	"net/url"
	"strconv"
	"strings"
	"time"
)

func createHandlerFunc(w http.ResponseWriter, r *http.Request) {
	suffix := strings.TrimSpace(r.FormValue("suffix"))
	target := strings.TrimSpace(r.FormValue("target"))
	update := strings.TrimSpace(r.FormValue("update")) == "true"
	ruleIdStr := strings.TrimSpace(r.FormValue("id"))
	var ruleId int64
	if update {
		if ruleIdStr == "" {
			util.Res{Message: "id 不能为空"}.Write(w)
			return
		}
		if id, err := strconv.ParseInt(ruleIdStr, 10, 64); err != nil {
			util.Res{Message: "id 格式错误，请输入 int"}.Write(w)
			return
		} else {
			ruleId = id
		}
	}

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
	rule := Rule{
		Suffix: suffix,
		Target: target,
		UserId: token.UserId,
	}
	if expiresString != "" {
		t, err := time.Parse(time.RFC3339, expiresString)
		if err != nil {
			util.Res{Message: "expires 字段必须是 RFC3339 格式的日期时间"}.Write(w)
			return
		}
		rule.Expires = &t
	}
	db := util.GetDB()
	defer db.Close()
	if update {
		rule.Id = ruleId
		err = rule.Update(db)
		if err != nil && err.Error() != "受影响的行数为 0" {
			util.Res{Message: "操作失败"}.Write(w)
		} else {
			util.Res{Success: true, Message: "更新成功", Data: rule}.Write(w)
		}

	} else {
		err = rule.Insert(db)
		if err != nil {
			if strings.Contains(err.Error(), "Duplicate entry") {
				util.Res{Message: "后缀已被使用，请换一个重试"}.Write(w)
			} else {
				util.Res{Message: "操作失败"}.Write(w)
			}
			return
		}
		util.Res{Success: true, Message: "创建成功", Data: rule}.Write(w)
	}
}
