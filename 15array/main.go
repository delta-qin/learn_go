package main

import "fmt"

func main() {
	// 初始化1
	var arr1 [5]int
	arr1 = [5]int{1, 0, 2, 2, 3}
	fmt.Println(arr1)

	// 初始化2
	arr2 := [5]int{1, 1, 1, 1, 1}
	fmt.Println(arr2)

	// 初始化3：不指定个数初始化
	arr3 := [...]int{2, 2, 2, 2, 2}
	fmt.Println(arr3)

	// 初始化4：按照位置初始化
	arr4 := [5]int{0: 1, 1: 2, 4: 6}
	fmt.Println(arr4)

	// 遍历1
	arr5 := [...]string{"1", "2", "3"}
	for i := 0; i < len(arr5); i++ {
		fmt.Println(arr5[i])
	}

	// 遍历2
	for i, v := range arr5 {
		fmt.Println(i, v)
	}

	// 二维数组
	var arr6 [3][2]int
	arr6 = [3][2]int{
		[2]int{1, 2},
		[2]int{3, 2},
		[2]int{5, 2},
	}

	for _, val1 := range arr6 {
		// fmt.Println(val1)
		for _, val2 := range val1 {
			fmt.Println(val2)
		}
	}

	// 数组是值类型，不同于引用类型
	arr7 := [2]int{1, 2}
	arr8 := arr7
	arr8[0] = 100
	fmt.Println(arr7)
	// [1 2]
	// [100 2]
	fmt.Println(arr8)
}
