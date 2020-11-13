package main

import "fmt"

// 只规定方法
type speaker interface {
	// 实现了speak方法的变量就是该接口类型
	speak()
}

type dog struct{}
type cat struct{}

func (d dog) speak() {
	fmt.Println("wang")
}

func (c cat)speak()  {
	fmt.Println("miao")
}

func attack(s speaker)  {
	s.speak()
}

func main() {
	var d dog
	var c cat
	attack(d)
	attack(c)
}
