package main

import "fmt"

func main() {
	data := Data{}
	data.Num = 9
	data.M()
	fmt.Println(data.Num)
}

func (data *Data) M() {
	data.Num = 123
}

type Data struct {
	Num int
}
