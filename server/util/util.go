package util

import (
	"log"
	"os"
)

// 检查环境变量完整性，要求 requires 中每个字段都不能为空，否则结束程序。
func RequireEnvs(requires ...string) {
	for _, key := range requires {
		if os.Getenv(key) == "" {
			log.Fatalln("环境变量缺少 " + key + " 字段")
		}
	}
}
