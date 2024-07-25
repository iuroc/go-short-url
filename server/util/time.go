package util

import (
	"log"
	"os"
	"time"
)

func GetLocation() *time.Location {
	timeZone := os.Getenv("DB_TIME_ZONE")
	loc, err := time.Parse("-07:00", timeZone)
	if err != nil {
		log.Fatalln("[GetLocation]", err)
	}
	offset := int(time.Date(0, 1, 1, 0, 0, 0, 0, time.UTC).Sub(loc).Seconds())
	location := time.FixedZone("", offset)
	return location
}

// ParseTimeFromDB 根据环境变量配置的 DB_TIME_ZONE 将从数据库中读取到的 Datetime 类型的时间转换为本地时间
func ParseTimeFromDB(datetime string) time.Time {
	t, err := time.ParseInLocation(time.DateTime, datetime, GetLocation())
	if err != nil {
		log.Fatalln("[ParseTimeFromDB]", err)
	}
	return t
}

// 将 time.Time 对象转换为不含时区信息的 Datetime 格式，
// 将根据 DB_TIME_ZONE 环境变量，生成对应时区的 datetime。
//
// 从前端读取 ISO 时间后，先转换为 [time.Time]，然后再用本方法转为 datetime 存入数据库。
func TimeToDatetime(t time.Time) string {
	return t.In(GetLocation()).Format(time.DateTime)
}