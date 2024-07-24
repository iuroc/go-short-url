package rule

import (
	"database/sql"
	"go-short-url/util"
	"log"
	"time"
)

type Rule struct {
	Id int64 `json:"id"`
	// 规则后缀，例如 example，则生成的短链接为 http://server/example
	Suffix string `json:"suffix"`
	// 重定向到的目标 URL
	Target string `json:"target"`
	// 访问次数统计
	Request int `json:"request"`
	// 创建者 ID
	UserId int64 `json:"userId"`
	// 链接过期时间，留空则永不过期
	Expires    *time.Time `json:"expires"`
	CreateTime time.Time  `json:"createTime"`
	UpdateTime time.Time  `json:"updateTime"`
}

func (rule Rule) Insert(db *sql.DB) (r *Rule, err error) {
	stmt, err := db.Prepare("INSERT INTO `go_short_url_rule` (`suffix`, `target`, `user_id`, `expires`) VALUES (?, ?, ?, ?)")
	if err != nil {
		log.Fatalln("[(rule Rule) Insert-1]", err)
	}
	var expiresString *string
	if rule.Expires != nil {
		s := rule.Expires.Format(time.DateTime)
		expiresString = &s
	}
	insertId, err := util.ExecErrorHandler(stmt.Exec(rule.Suffix, rule.Target, rule.UserId, expiresString))
	if err != nil {
		return nil, err
	} else {
		rule.Id = insertId
		rule.CreateTime = time.Now()
		rule.UpdateTime = time.Now()
		return &rule, nil
	}
}

func SelectTargetBySuffix(db *sql.DB, suffix string) (*Rule, error) {
	stmt, err := db.Prepare("SELECT `id`, `suffix`, `target`, `request`, `user_id`, `expires`, `create_time`, `update_time` FROM `go_short_url_rule` WHERE `suffix` = ?")
	if err != nil {
		return nil, err
	}
	rule := Rule{}
	row := stmt.QueryRow(suffix)
	createTime := ""
	updateTime := ""
	expires := sql.NullString{}
	err = row.Scan(&rule.Id, &rule.Suffix, &rule.Target, &rule.Request, &rule.UserId, &expires, &createTime, &updateTime)
	if err != nil {
		log.Fatalln(err)
	}
	rule.CreateTime = util.ParseTimeFromDB(createTime)
	rule.UpdateTime = util.ParseTimeFromDB(updateTime)
	if expires.Valid {
		datetime := util.ParseTimeFromDB(expires.String)
		rule.Expires = &datetime
	}
	return &rule, nil
}
