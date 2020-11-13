package main

import "fmt"

type dog struct {
	name string
}

func newDog(name string ) *dog {
	return &dog{
		name: name,
	}
}
// 和函数的区别就是函数名前面需要制定具体的类型变量，
// 多用类型名首字母小写
func (d dog) wang() {
	// fmt.Println("wangwang")
	d.name = "1"
}

func (d *dog) wang1() {
	// fmt.Println("wangwang")
	d.name = "2"
}

type myInt int

func (i myInt) my()  {
	fmt.Println("直接操作int不可以，不是自己包定义的")
}

func main() {
	// 方法是作用于特定类型的函数
	d1 := newDog("de")
	d1.wang()
	fmt.Println(d1)
	d1.wang1()
	fmt.Println(d1)

	m := myInt(10)
	m.my()
	fmt.Println(m)
}
