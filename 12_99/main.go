package main

import "fmt"

func main() {
	for i := 1; i <= 9; i++ {
		// 不写=会每一行缺一项
		for j := 1; j <= i; j++ {
			fmt.Printf("%d*%d=%d\t", j, i, i*j)
		}
		fmt.Println()
	}
}