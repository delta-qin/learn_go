package main

import "fmt"

type st1 struct {
	name string
	age  int8
}

func f(x st1) {
	x.name = "qin1"
}

func f1(x *st1) {
	(*x).name = "qin1"
}

func f2(x *st1) {
	x.name = "qin2"
}

func main() {
	var p st1
	p.name = "qin"
	p.age = 1
	f(p)
	fmt.Println(p)
	f1(&p)
	fmt.Println(p)
	f2(&p)
	fmt.Println(p)
	fmt.Printf("%p\n", &p) // 0xc0000044a0

	// p1 是结构体指针
	var p1 = new(st1)
	fmt.Printf("%x\n", p1) // &{ 0}
	fmt.Printf("%x\n", &p1) // c000006030
	fmt.Printf("%p\n", &p1) // 0xc000006030
	fmt.Printf("%#v\n", p1) // &main.st1{name:"", age:0}
	fmt.Printf("%T\n", p1) // *main.st1
	// 可以直接对结构体指针使用 . 来获取属性
	p1.name = "deltaqin"
	p1.age = 12
	fmt.Printf("%#v\n", p1) // &main.st1{name:"deltaqin", age:12}


	p2 := st1{
		name: string("deltaqin"),
		age: int8(1),
	}

	fmt.Printf("%p\n", &(p2.name))
	fmt.Printf("%p\n", &(p2.age))
	// 0xc0000045c0，16位
	// 0xc0000045d0，16位
	// 还有内存对齐，在必要的时候
}
