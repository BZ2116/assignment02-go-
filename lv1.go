package main

import "fmt"

type player struct {
	Name   string
	Height int
	Weight int
	Age    int8
	number int
}

func main() {

	Rose := player{
		Name:   "Rose",
		Height: 191,
		Weight: 90,
		Age:    34,
		number: 1,
	}
	fmt.Printf("姓名:%s,身高：%d，体重：%d，年龄：%d,球衣号码:%d\n", Rose.Name, Rose.Height, Rose.Weight, Rose.Age, Rose.number)
}
