package util

import (
	"database/sql"
	"errors"
	"log"
)

// 处理执行 Exec 后的异常，返回 LastInsertId 和需要抛出的错误
func HandleExecError(result sql.Result, err error) (int64, error) {
	if err != nil {
		return 0, errors.New("操作失败，" + err.Error())
	}
	if rows, err := result.RowsAffected(); err != nil {
		log.Fatalln(err)
	} else if rows == 0 {
		return 0, errors.New("操作失败，受影响的行数为 0")
	}
	id, err := result.LastInsertId()
	if err != nil {
		log.Fatalln(err)
	}
	return id, nil
}
