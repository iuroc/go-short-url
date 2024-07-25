// 用户相关的路由
//
// 实现登录注册功能
package user

import (
	"database/sql"
	"go-short-url/util"
	"golang.org/x/crypto/bcrypt"
	"log"
	"os"
	"strings"
	"time"
)

// User 用户信息
//
// JWT 需要携带 Id 和 Username
type User struct {
	Id           int64     `json:"id"`
	Username     string    `json:"username"`
	PasswordHash string    `json:"-"`
	Role         string    `json:"-"`
	CreateTime   time.Time `json:"createTime"`
	UpdateTime   time.Time `json:"updateTime"`
}

func (u *User) Insert(db *sql.DB) error {
	stmt, err := db.Prepare("INSERT INTO `go_short_url_user` (`username`, `password`, `role`) VALUES (?, ?, ?)")
	if err != nil {
		log.Fatalln("[(u User) Insert-1]", err)
	}
	if insertId, err := util.ExecErrorHandler(stmt.Exec(u.Username, u.PasswordHash, u.Role)); err != nil {
		return err
	} else {
		u.Id = insertId
		u.CreateTime = time.Now()
		u.UpdateTime = time.Now()
		return nil
	}
}

func InitAdmin(db *sql.DB) {
	username := strings.TrimSpace(os.Getenv("ROOT_USERNAME"))
	password := strings.TrimSpace(os.Getenv("ROOT_PASSWORD"))
	if err := util.CheckUsernameAndPasswordFormat(username, password); err != nil {
		log.Fatalln("[InitAdmin-1]", err)
	}
	if user := SelectUserByRole(db, "admin"); user != nil {
		log.Println("管理员账户已存在，无需操作")
		return
	}
	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(password), bcrypt.DefaultCost)
	if err != nil {
		log.Fatalln("[InitAdmin-2]", err)
	}
	user := User{
		Username:     username,
		PasswordHash: string(hashedPassword),
		Role:         "admin",
	}
	if err = user.Insert(db); err != nil {
		log.Fatalln("[InitAdmin-3]", err)
	}
	log.Println("管理员账户初始创建完成")
}

// SelectUserByName 通过用户名查找用户记录
//
// fields 示例：
//
//	[][]string{
//		{"username", "zhangsan"},
//		{"password", "12345678"},
//		{"role", "user"}
//	}
func SelectUserByFields(db *sql.DB, fields map[string]any) *User {
	query := "SELECT `id`, `username`, `password`, `role`, `create_time`, `update_time` FROM `go_short_url_user` WHERE "
	keyTexts := []string{}
	values := []any{}
	for key, value := range fields {
		keyTexts = append(keyTexts, key+" = ?")
		values = append(values, value)
	}
	query = query + strings.Join(keyTexts, " AND ")
	stmt, err := db.Prepare(query)
	if err != nil {
		log.Fatalln("[SelectUserByFields-1]", err)
	}
	user := User{}
	row := stmt.QueryRow(values...)
	var createTime string
	var updateTime string
	err = row.Scan(&user.Id, &user.Username, &user.PasswordHash, &user.Role, &createTime, &updateTime)
	if err == sql.ErrNoRows {
		return nil
	} else if err != nil {
		log.Fatalln("[SelectUserByFields-2]", err)
	}
	user.CreateTime = util.ParseTimeFromDB(createTime)
	user.UpdateTime = util.ParseTimeFromDB(updateTime)
	return &user
}

func SelectUserByRole(db *sql.DB, role string) *User {
	return SelectUserByFields(db, map[string]any{
		"role": role,
	})
}

func SelectUserByName(db *sql.DB, username string) *User {
	return SelectUserByFields(db, map[string]any{
		"username": username,
	})
}

func SelectUserByNameAndPassword(db *sql.DB, username string, password string) *User {
	return SelectUserByFields(db, map[string]any{
		"username": username,
		"password": username,
	})
}
