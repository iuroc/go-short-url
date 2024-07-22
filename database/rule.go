package database

import "time"

// 短网址跳转规则
type Rule struct {
	// 记录 ID
	Id int64 `json:"id"`
	// 规则后缀
	Suffix string `json:"suffix"`
	// 目标 URL
	Target string `json:"target"`
	// 访问次数统计
	Request int `json:"request"`
	// 创建者 ID
	UserId int `json:"userId"`
	// 过期时间
	Expires    time.Time `json:"expires"`
	CreateTime time.Time `json:"createTime"`
	UpdateTime time.Time `json:"updateTime"`
}
