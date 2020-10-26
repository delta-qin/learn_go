package main

import "fmt"

func main() {
	var (
		name string
		age  int
	)

	// 没有引号
	fmt.Scan(&name, &age)
	fmt.Println(name, age)
	fmt.Scanln(&name, &age)
	fmt.Println(name, age)
	fmt.Scanf("%s %d\n", &name, &age)
	fmt.Println(name, age)
}
