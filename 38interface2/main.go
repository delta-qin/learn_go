package main

import "fmt"

type animal interface {
	eat()
}

type cat struct{}
type dog struct{}

func (c cat) eat() {
	fmt.Println("fish")
}

func (d dog) eat() {
	fmt.Println("骨头")
}

func main() {
	var a animal
	a = cat{}
	a.eat()
	a = dog{}
	a.eat()
}
