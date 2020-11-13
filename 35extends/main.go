package main

import "fmt"

type animal struct {
	name string
}

func (a animal) move() {
	fmt.Printf("%s move", a.name)
}

type dog struct {
	feet uint8
	animal
}

func (d dog) wang()  {
	fmt.Printf("%s wang\n", d.name)
}
func main() {
	d := dog{
		feet:4,
		animal: animal{
			name: "de",
		},
	}
	d.wang()
	// 继承匿名 animal，方法就全有了，变相实现继承
	d.move()
}
