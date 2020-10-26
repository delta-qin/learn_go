package main

import "fmt"

func f1() {
	// 如果是在defer后面使用匿名函数，一定要执行，不要忘记立即执行
	defer func() {
		err := recover()
		fmt.Println(err)
	}()
	panic("严重错误")
	fmt.Println("这里不会执行")
}

func main() {
	f1()
	fmt.Println("这里会执行")
}
