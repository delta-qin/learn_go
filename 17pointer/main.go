package main

import "fmt"

func main() {

	// & *
	n := 18
	p := &n
	t := *p

	fmt.Println(p)
	fmt.Println(t)
	fmt.Printf("%T\n", p)
	
	fmt.Println("===============")

	// new 返回的是空间的地址
	var p1 *int
	fmt.Println(p1)
	p1 = new(int)
	fmt.Println(p1)
	fmt.Println(*p1)
	*p1 = 100
	fmt.Println(*p1)

	// make
	

}
