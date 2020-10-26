package main

import "fmt"

func f1(f func()) {
	fmt.Println("f1")
	f()
}

func f2(x, y int) {
	fmt.Println("f2")
	fmt.Println(x, y)
}

func f3(f func(int, int), m, n int) func() {
	// 闭包：一个函数可以引用他外面的变量，这里的f和mn都是环境的变量
	// 最后main调用的时候看起来是使用了 temp ，实际上就是装饰者模式，
	// 在原有的方法基础上定义自己的相关属性
	temp := func() {
		f(m, n)
	}
	return temp
}

func main() {
	// 匿名函数的使用：闭包
	// 立即执行，相当于调用了 f2
	f3(f2, 1, 2)()
}
