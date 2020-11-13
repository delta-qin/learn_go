package main

import "fmt"

// 丰富维度的数据
// 自己造类型
type st1 struct {
	name  string
	age   int
	gen   string
	hobby []string
}

func main() {
	var de st1
	de.name = "de"
	de.age = 12
	de.gen = "男"
	de.hobby = []string{"篮球", "足球"}

	fmt.Println(de)

	fmt.Printf("%T\n", de)

	fmt.Println(de.age)

	// 匿名结构体
	var s struct{
		x string
		y int 
	}

	s.x = "de"
	fmt.Printf("%T\n", s)
	fmt.Println(s)

	// 指针类型结构体
	
}
