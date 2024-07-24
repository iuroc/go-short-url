package main

import (
	"fmt"
	"log"
	"time"
)

func main() {

	// 数据库里面的 Datetime 字符串
	datetime := "2024-07-24 13:39:30"
	// 数据库的时区
	timeZone := "+08:00"
	tt, _ := time.Parse("-07:00", timeZone)
	loc := time.FixedZone("xxx", int(time.Date(0, 1, 1, 0, 0, 0, 0, time.UTC).Sub(tt).Seconds()))
	t, err := time.ParseInLocation(time.DateTime, datetime, loc)
	if err != nil {
		log.Fatalln(err)
	}
	fmt.Println(t)
	fmt.Println(t.UTC())
}
