package main

import "fmt"

func main() {

	// 切片和map使用的时候一定要初始化，
	// 可以使用make，也可以直接赋值初始化

	// 值为map的切片
	// 注意使用make分配切片内存空间的时候，size一定要指定好，
	// 否则size为0，直接使用 = 赋值会失败，只能使用append，因为切片窗口是0
	s1 := make([]map[string]int, 10, 10)
	s1[0] = make(map[string]int,1)
	s1[0]["deltaqin"] = 20
	fmt.Println(s1)

	// 值为切片的map
	// map 在分配map的时候只需要cap，不要多写参数
	m1 := make(map[string][]int,10)
	m1["deltaqin"] = []int{1,2,3}
	fmt.Println(m1)

	// 输出
	// [map[deltaqin:20] map[] map[] map[] map[] map[] map[] map[] map[] map[]]
	// map[deltaqin:[1 2 3]]
}
