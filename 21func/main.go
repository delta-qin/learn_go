package main

import "fmt"

func f2(x int, y int) {
	fmt.Println("没有返回值")
}

func f3() {
	fmt.Println("没有参数没有返回值")
}

func f4() int {
	fmt.Println("没有参数有返回值")
	ret := 3
	return ret
}

func f5(x int, y int) (ret int) {
	fmt.Println("命名的返回值，相当于在函数里面声明了一个变量")
	ret = x + y
	return // 命名返回值，return后面可以省略
}

func f6() (int, string) {
	fmt.Println("多个返回值")
	return 1, "deltaqin"
}

func f7(x, y int, z, m string) int {
	fmt.Println("参数类型简写，连续多个参数类型一致，可以将最后一个之前的参数类型去掉")
	return x + y
}

func f8(x string, y ...int) {
	fmt.Println("可变长参数，必须放在最后")
	fmt.Println(x)
	fmt.Println(y)
	// y是切片类型 []int
}

func main() {

	f2(1, 2)
	f3()
	r4 := f4()
	r5 := f5(1, 2)
	r6, r61 := f6()
	r7 := f7(1, 2, "1", "2")
	f8("1",1,2,3)
	fmt.Println(r4, r5, r6, r61, r7)

}
