package main

import "fmt"

func main() {

	// 1 匿名函数赋值给变量
	var anonymous = func(x, y int) {
		fmt.Println(x, y)
	}
	anonymous(1,2)

	// 2 立即执行函数
	func (x,y int)  {
		fmt.Println(x, y)
	}(1,2)

}
