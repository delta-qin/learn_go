package main

import "fmt"

type person struct {
	name string
	age  int
}

func newPerson(name string, age int) *person {
	return &person{
		name: name,
		age:  age,
	}
}

func newPerson2(name string, age int) person {
	return person{
		name: name,
		age:  age,
	}
}

func main() {
	p1 := newPerson("deltaqin", 18)
	fmt.Println(p1)
	fmt.Printf("%p\n", &p1)
	// &{deltaqin 18}
	// 0xc000006028
	p2 := newPerson2("deltaqin", 18)
	fmt.Println(p2)
	fmt.Printf("%p\n", &p2)
	// {deltaqin 18}
	// 0xc000004500
}
