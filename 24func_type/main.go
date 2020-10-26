package main

import "fmt"

func f1() {
	fmt.Println("1")
}

func f2() int {
	fmt.Println("2")
	return 10
}

func f3(x func() int) {
	ret := x()
	fmt.Println(ret)
}

func f4(x, y int) int {
	return x + y
}

func temp(a, b int) int {
	return a + b
}

func f5(x func() int) func(int,int) int {
	return temp
}

func main() {
	a := f1
	fmt.Printf("%T\n", a)

	b := f2
	fmt.Printf("%T\n",b)

	f3(f2)

	d := f4
	fmt.Printf("%T\n",d)

	// 注意传递参数的时候，函数的类型要一致（包括参数以及返回值的个数和类型）
	e := f5(f2)
	fmt.Printf("%T\n", e)

	
	// func()
	// func() int
	// 2
	// 10
	// func(int, int) int
	// func(int, int) int
}
