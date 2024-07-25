package rule

import (
	"database/sql"
	"go-short-url/util"
	"log"
	"strings"
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

func (rule *Rule) Insert(db *sql.DB) error {
	stmt, err := db.Prepare("INSERT INTO `go_short_url_rule` (`suffix`, `target`, `user_id`, `expires`) VALUES (?, ?, ?, ?)")
	if err != nil {
		log.Fatalln("[(rule Rule) Insert-1]", err)
	}
	var expiresString sql.NullString
	if rule.Expires != nil {
		expiresString.String = rule.Expires.In(util.GetLocation()).Format(time.DateTime)
		expiresString.Valid = true
	}
	insertId, err := util.ExecErrorHandler(stmt.Exec(rule.Suffix, rule.Target, rule.UserId, expiresString))
	if err != nil {
		return err
	} else {
		rule.Id = insertId
		rule.CreateTime = time.Now()
		rule.UpdateTime = time.Now()
		return nil
	}
}

func (rule *Rule) Update(db *sql.DB) error {
	stmt, err := db.Prepare("UPDATE `go_short_url_rule` SET `target` = ?, `request` = ?, `user_id` = ?, `expires` = ? WHERE `id` = ?")
	if err != nil {
		return err
	}
	var expires sql.NullTime
	if rule.Expires != nil {
		expires.Valid = true
		expires.Time = rule.Expires.In(util.GetLocation())
	}
	_, err = util.ExecErrorHandler(stmt.Exec(rule.Target, rule.Request, rule.UserId, expires, rule.Id))
	if err != nil {
		return err
	}
	return nil
}

const SELECT_PREFIX = "SELECT `id`, `suffix`, `target`, `request`, `user_id`, `expires`, `create_time`, `update_time` FROM `go_short_url_rule`"

func SelectRuleBySuffix(db *sql.DB, suffix string) (*Rule, error) {
	stmt, err := db.Prepare(SELECT_PREFIX + " WHERE `suffix` = ?")
	if err != nil {
		return nil, err
	}
	rule := Rule{}
	row := stmt.QueryRow(suffix)
	createTime := ""
	updateTime := ""
	expires := sql.NullString{}
	err = row.Scan(&rule.Id, &rule.Suffix, &rule.Target, &rule.Request, &rule.UserId, &expires, &createTime, &updateTime)
	if err == sql.ErrNoRows {
		return nil, err
	}
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

func DeleteById(db *sql.DB, userId int64, id int64) error {
	stmt, err := db.Prepare("DELETE FROM `go_short_url_rule` WHERE `id` = ? AND `user_id` = ?")
	if err != nil {
		log.Fatalln("[DeleteById]", err)
	}
	_, err = util.ExecErrorHandler(stmt.Exec(id, userId))
	if err != nil {
		return err
	}
	return nil
}

func SearchRules(db *sql.DB, userId int64, keyword string, page int, pageSize int) ([]Rule, error) {
	stmt, err := db.Prepare(SELECT_PREFIX + " WHERE CONCAT(`suffix`, `target`) LIKE ? AND `user_id` = ? LIMIT ? OFFSET ?")
	if err != nil {
		log.Fatalln("[SearchRules-1]", err)
	}
	rows, err := stmt.Query("%"+strings.Join(strings.Split(keyword, " "), "%")+"%", userId, pageSize, page*pageSize)
	if err != nil {
		return nil, err
	}
	var rules []Rule
	defer rows.Close()
	for rows.Next() {
		rule := Rule{}
		createTime := ""
		updateTime := ""
		expires := sql.NullString{}
		err = rows.Scan(&rule.Id, &rule.Suffix, &rule.Target, &rule.Request, &rule.UserId, &expires, &createTime, &updateTime)
		if err != nil {
			log.Fatalln(err)
		}
		rule.CreateTime = util.ParseTimeFromDB(createTime)
		rule.UpdateTime = util.ParseTimeFromDB(updateTime)
		if expires.Valid {
			datetime := util.ParseTimeFromDB(expires.String)
			rule.Expires = &datetime
		}
		rules = append(rules, rule)
	}
	return rules, nil
}
