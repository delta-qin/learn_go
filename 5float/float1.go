package main

import "fmt"

// 浮点数

func main() {
	// math.MaxFloat32
	f1 := 1.2333 // 默认64，32需要显式说明
	fmt.Printf("%T\n", f1) //float64
	f2 := float32(1.223)
	fmt.Printf("%T\n", f2) //float32 

	// 不能直接赋值，即使是小转大

}
