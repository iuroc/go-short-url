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
	
	insertId, err := util.ExecErrorHandler(stmt.Exec(rule.Suffix, rule.Target, rule.UserId, rule.Expires))
	if err != nil {
		return nil, err
	} else {
		rule.Id = insertId
		rule.CreateTime = time.Now()
		rule.UpdateTime = time.Now()
		return &rule, nil
	}
}
