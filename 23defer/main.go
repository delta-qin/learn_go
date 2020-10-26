package main

import "fmt"

func f1() int {
	x := 5
	defer func() {
		x++
	}()
	return x // 5 x是函数自己定义的变量，不是返回变量
}

func f2() (x int) {
	defer func() {
		x++
	}()
	return 5 // 6 改变的就是返回值x，defer自己没有声明变量
}

func f3() (y int) {
	x := 5
	defer func() {
		x++
	}()
	return x // 5 没有改变y
}

func f4() (x int) {
	defer func(x int) {
		x++
	}(x)
	return 5 // 5 函数传参是拷贝
}

func f5() (x int) {
	defer func(x *int) {
		(*x)++
	}(&x)
	return 5 // 6 立即执行函数传递进去的是指针
}

func calc(index string, a, b int) int {
	ret := a + b
	fmt.Println(index, a, b, ret)
	return ret
}

func main() {
	// // defer 延迟执行，当前函数即将返回的时候再执行
	// fmt.Println("start")
	// defer fmt.Println(1)
	// defer fmt.Println(2)
	// defer fmt.Println(3)
	// fmt.Println("end")

	// // start
	// // end
	// // 3
	// // 2
	// // 1

	r1 := f1()
	r2 := f2()
	r3 := f3()
	r4 := f4()
	fmt.Println(r1, r2, r3, r4)

	a := 1
	b := 1
	// 执行到这里会先将参数里面的 calc("10", a, b) 算出来
	// 再把calc整体压入栈里面，并不是到时候一起算
	defer calc("1", a, calc("10", a, b))
	a = 0
	// 执行到这里会先将参数里面的 calc("20", a, b) 算出来
	// 再把calc整体压入栈里面
	defer calc("2", a, calc("20", a, b))
	b = 1
	fmt.Println(111)
	// 10 1 1 2
	// 20 0 1 1
	// 111
	// 2 0 1 1
	// 1 1 2 3
}
