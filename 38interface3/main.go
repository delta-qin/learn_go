package main

import "fmt"

type animal interface {
	eat()
}

type cat struct{}
type dog struct{}

// func (c cat) eat() {
// 	fmt.Println("fish")
// }

// func (d dog) eat() {
// 	fmt.Println("骨头")
// }

func (c *cat) eat() {
	fmt.Println("fish")
}

func (d *dog) eat() {
	fmt.Println("骨头")
}

func main() {
	// var a1 animal
	// c1 := cat{}
	// a1 = c1
	// a1.eat()

	var a1 animal
	c1 := cat{}
	a1 = &c1
	a1.eat()
}