package database

import (
	"database/sql"
	"go-short-url/util"
	"log"
	"strings"
)

type User struct {
	// 记录编号
	Id int64 `json:"id"`
	// 用户名
	Username string `json:"username"`
	// bcrypt 哈希密码
	Password string `json:"-"`
	// 用户身份，admin 管理员，user 普通用户
	Role string `json:"-"`
	// 创建时间
	CreateTime string `json:"create_time"`
}

func (u User) Insert(db *sql.DB) (user *User, err error) {
	stmt, err := db.Prepare("INSERT INTO `go_short_url_user` (`username`, `password`, `role`) VALUES (?, ?, ?)")
	if err != nil {
		log.Fatalln("[Insert]", err)
	}
	if lastId, err := util.HandleExecError(
		stmt.Exec(u.Username, u.Password, u.Role),
	); err != nil {
		// 插入失败，用户原因的错误
		return nil, err
	} else {
		// 插入成功
		u.Id = lastId
		u.CreateTime = util.GetNowDatetimeString()
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
		log.Fatalln(err)
	}
	_, err = util.HandleExecError(stmt.Exec(values...))
	return err
}

// 验证账号密码，密码是明文密码
func CheckLogin(db *sql.DB, username string, password string) bool {
	hashedPassword, err := GetUserHashedPassword(db, username)
	if err != nil {
		return false
	}
	return util.CheckPasswordHash(password, hashedPassword)
}

// 从数据库获取用户的哈希密码，如果没有找到，则返回 error
func GetUserHashedPassword(db *sql.DB, username string) (string, error) {
	stmt, err := db.Prepare("SELECT `password` FROM `go_short_url_user` WHERE `username` = ?")
	if err != nil {
		log.Fatalln(err)
	}
	row := stmt.QueryRow(username)
	var hashedPassword string
	err = row.Scan(&hashedPassword)
	if err != nil {
		return "", err
	}
	return hashedPassword, nil
}
