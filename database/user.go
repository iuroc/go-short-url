package database

import (
	"database/sql"
	"go-short-url/mixin"
	"log"
	"strings"
	"time"
)

type User struct {
	// 记录 ID
	Id int64 `json:"id"`
	// 用户名
	Username string `json:"username"`
	// bcrypt 哈希密码
	Password string `json:"-"`
	// 用户身份，admin 管理员，user 普通用户
	Role string `json:"-"`
	// 创建时间
	CreateTime time.Time `json:"createTime"`
	UpdateTime time.Time `json:"updateTime"`
}

func (u User) Insert(db *sql.DB) (user *User, err error) {
	stmt, err := db.Prepare("INSERT INTO `go_short_url_user` (`username`, `password`, `role`) VALUES (?, ?, ?)")
	if err != nil {
		log.Fatalln("[Insert]", err)
	}
	if lastId, err := mixin.HandleExecError(
		stmt.Exec(u.Username, u.Password, u.Role),
	); err != nil {
		// 插入失败，用户原因的错误
		return nil, err
	} else {
		// 插入成功
		u.Id = lastId
		u.CreateTime = time.Now()
		u.UpdateTime = time.Now()
		return &u, nil
	}
}

// 判断管理员账户是否存在
func AdminExists(db *sql.DB) (bool, int64) {
	stmt, err := db.Prepare("SELECT `id` FROM `go_short_url_user` WHERE `role` = 'admin'")
	if err != nil {
		log.Fatalln("[AdminExists-1]", err)
	}
	row := stmt.QueryRow()
	var id int64
	err = row.Scan(&id)
	if err == sql.ErrNoRows {
		return false, 0
	} else if err != nil {
		log.Fatalln("[AdminExists-2]", err)
	}
	return true, id
}

// 只更新账号或密码
func (u User) Update(db *sql.DB) error {
	fields := make([]string, 0)
	values := make([]any, 0)
	if u.Username != "" {
		fields = append(fields, "`username` = ?")
		values = append(values, u.Username)
	}
	if u.Password != "" {
		fields = append(fields, "`password` = ?")
		values = append(values, u.Password)
	}
	values = append(values, u.Id)
	stmt, err := db.Prepare("UPDATE `go_short_url_user` SET " + strings.Join(fields, ", ") + " WHERE `id` = ?")
	if err != nil {
		log.Fatalln("[(u User) Update]", err)
	}
	_, err = mixin.HandleExecError(stmt.Exec(values...))
	return err
}

// 验证账号密码，密码是明文密码，如果检查通过，则会返回 User 对象，否则返回 nil
func CheckLogin(db *sql.DB, username string, password string) *User {
	user, err := GetUserByUsername(db, username)
	hashedPassword := user.Password
	if err != nil {
		return nil
	}
	if mixin.CheckPasswordHash(password, hashedPassword) {
		return user
	} else {
		return nil
	}
}

// 根据用户名获取 User 对象
func GetUserByUsername(db *sql.DB, username string) (*User, error) {
	stmt, err := db.Prepare("SELECT `id`, `username`, `password`, `role`, `create_time`, `update_time` FROM `go_short_url_user` WHERE `username` = ?")
	if err != nil {
		log.Fatalln("[GetUserByUsername-1]", err)
	}
	row := stmt.QueryRow(username)
	user := User{}
	err = row.Scan(&user.Id, &user.Username, &user.Password, &user.Role, &user.CreateTime, &user.UpdateTime)
	if err == sql.ErrNoRows {
		return nil, err
	} else if err != nil {
		log.Fatalln("[GetUserByUsername-2]", err)
	}
	return &user, nil
}
