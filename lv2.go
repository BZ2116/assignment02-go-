package main

import "fmt"

type Player interface {
	number()
	team()
}

type Rose struct {
}
type Leborn struct {
}
type Yao struct {
}

func (a Rose) number() {
	fmt.Println("球衣号码：1号")
}
func (a Rose) team() {
	fmt.Println("球队：公牛队")
}
func (a Leborn) number() {
	fmt.Println("球衣号码：23号")
}
func (a Leborn) team() {
	fmt.Println("球队：骑士队")
}

func (a Yao) number() {
	fmt.Println("球衣号码：11号")
}
func (a Yao) team() {
	fmt.Println("球队：火箭队")
}
func show(my Player) {
	my.number()
	my.team()
}

func main() {
	show(new(Rose))
}
