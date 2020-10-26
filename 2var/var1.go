package main

import "fmt"

// const name = "qzt"
// const age = 12
// const flag = false

const (
	name = "qzt"
	age  = 12
	flag = false
)

const (
	n1 = iota // 0
	n2        // 1
	n3        // 2
)

const (
	m1 = iota // 0
	m2        //1
	_
	m3 //3
)

const (
	c1 = iota // 0
	c2 = 100
	c3 = iota // 2 新增一行就会加1
	c4        // 3
)

const (
	a1, a2 = iota + 1, iota + 2 // 1,2
	b1, b2 = iota + 1, iota + 2 // 2,3
)

// 定义数量级
const (
	_ = iota
	// 1左移10位相当于乘了2^10，1000000000=1024
	KB = 1 << (10 * iota)
	// 左移20位,2^20 = 2^10*2^10 = 1024*1024
	MB = 1 << (10 * iota)
	GB = 1 << (10 * iota)
	TB = 1 << (10 * iota)
	PB = 1 << (10 * iota)
)

func main() {
	fmt.Println("name：", name)
	fmt.Printf("age：%d\n", age)
	fmt.Print("flag：", flag)
	fmt.Println("===================")

	fmt.Println("n1", n1)
	fmt.Println("n2", n2)
	fmt.Println("n3", n3)

	fmt.Println("m1", m1)
	fmt.Println("m2", m2)
	fmt.Println("m3", m3)

	fmt.Println("=========插队==========")
	fmt.Println("c1", c1)
	fmt.Println("c2", c2)
	fmt.Println("c3", c3)
	fmt.Println("c4", c4)

	fmt.Println("=========一行多常量==========")
	fmt.Println("a1", a1)
	fmt.Println("a2", a2)
	fmt.Println("b1", b1)
	fmt.Println("b2", b2)
}
