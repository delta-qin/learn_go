package main

import "fmt"

func main() {
	var a1 = 100
	fmt.Printf("%d\n", a1) //100
	fmt.Printf("%b\n", a1) //1100100
	fmt.Printf("%o\n", a1) //144
	fmt.Printf("%x\n", a1) // 64

	a2 := 017
	fmt.Printf("%d\n", a2) // 15

	a3 := 0x123
	fmt.Printf("%d\n", a3) // 291

	a4 := int8(1)
	fmt.Printf("%d\n", a4) // 1

	// 查看变量类型
	fmt.Printf("%T\n", a1)
	fmt.Printf("%T\n", a4) //int8
}
