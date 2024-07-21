package database

import (
	"database/sql"
	"errors"
)

type User struct {
	// 记录编号
	Id int64 `json:"id"`
	// 用户名
	Username string `json:"username"`
	// bcrypt 哈希密码
	Password string
	// 创建时间
	CreateTime string `json:"create_time"`
}

func (u User) Insert(db *sql.DB) (user *User, err error) {
	if stmt, err := db.Prepare("INSERT INTO `go_short_url_user` (`username`, `password`) VALUES (?, ?)"); err != nil {
		return nil, err
	} else {
		if result, err := stmt.Exec(u.Username, u.Password); err != nil {
			return nil, err
		} else {
			if rows, err := result.RowsAffected(); err != nil {
				return nil, err
			} else if rows == 0 {
				return nil, errors.New("插入失败")
			}
			id, err := result.LastInsertId()
			if err != nil {
				return nil, err
			}
			u.Id = id
			u.CreateTime = "1234"
			return &u, nil
		}
	}
}
