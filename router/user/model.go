// 用户相关的路由
//
// 实现登录注册功能
package user

import "time"

// User 用户信息
//
// JWT 需要携带 Id 和 Username
type User struct {
	Id           int64     `json:"id"`
	Username     string    `json:"username"`
	PasswordHash string    `json:"-"`
	Role         string    `json:"-"`
	CreateTime   string    `json:"createTime"`
	UpdateTime   time.Time `json:"updateTime"`
}
