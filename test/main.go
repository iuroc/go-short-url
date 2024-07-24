package main

import (
	"fmt"
)

func main() {
	a := 0
	b := &a
	*b = 100
	fmt.Println(a)
}
