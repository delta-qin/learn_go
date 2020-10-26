package main

import "fmt"

func main() {
	bool1 := false
	fmt.Printf("%T\n", bool1) // bool

	var bool2 bool
	fmt.Printf("%T : %v", bool2, bool2) //bool : false
}
