package main

import "fmt"

func main() {
	age := 20
	if age > 18 {
		fmt.Println(1)
	} else if age < 10 {
		fmt.Println(0)
	} else {
		fmt.Println(2)
	}

	// 特殊写法
	// 局部变量，作用域，用完即销毁，后面获取不到
	if age1:=20; age1>19 {
		fmt.Println(99)
	} else {
		fmt.Println(22)
	}
}
