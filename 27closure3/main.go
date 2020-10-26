package main

import "fmt"

func call(base int) (func(int) int, func(int) int) {
	// 注意这里的base被修改之后，两个内部函数是相互影响的，不同调用次数之间也是相互影响的
	// 只要是只声明了一次 call 函数。改变的base就始终是那一个
	
	add := func(i int) int {
		base += i
		return base
	}

	sub := func(i int) int {
		base -= i
		return base
	}

	return add, sub
}

func main() {
	f1, f2 := call(10)

	fmt.Println(f1(1), f2(2))
	fmt.Println(f1(3), f2(4))
}
