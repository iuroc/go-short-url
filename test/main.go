package main

import (
	"fmt"
	"time"
)

func main() {
	t, err := time.Parse(time.RFC3339, "2024-07-20T19:29:05+08:00")
	fmt.Println(t.UTC(), err)
}
