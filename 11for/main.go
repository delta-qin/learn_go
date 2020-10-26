package main

import "fmt"

func main() {
	for i := 0; i < 10; i++ {
		fmt.Println(i)
	}

	// 1
	i:= 5
	for ; i < 10; i++ {
		fmt.Println(i)
	}

	// 2
	for i<10 {
		fmt.Println(i)
		i++
	}

	// 死循环，电脑差点起飞
	// for {
	// 	fmt.Println("1")
	// }
	

	// range
	s1 := "deltaqin你好"
	for i,v := range s1{
		// 直接打印是ASCII
		// fmt.Println(i,v)

		fmt.Printf("%d,%c\n" ,i,v)
	}
}
