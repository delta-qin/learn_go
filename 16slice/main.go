package main

import "fmt"

func main() {
	// 1. 定义切片：不指定个数
	var s1 []int
	var s2 []string
	fmt.Println(s1, s2)
	fmt.Println(s1 == nil) // true 没有开辟内存空间
	fmt.Println(s2 == nil)

	// 初始化切片
	s1 = []int{1, 2, 3}
	s2 = []string{"1", "2", "3"}
	fmt.Println(s1, s2)
	fmt.Println(s1 == nil) // false
	fmt.Println(s2 == nil)

	// 长度和容量
	fmt.Printf("len(s1):%d, cap(s1):%d\n", len(s1), cap(s1))
	fmt.Printf("len(s2):%d, cap(s2):%d\n", len(s2), cap(s2))

	// 2. 数组得到切片
	a1 := [...]int{1, 2, 3, 4, 5, 6}
	s3 := a1[0:4]   // [,)
	fmt.Println(s3) // [1 2 3 4]
	s4 := a1[:4]
	fmt.Println(s4) // [1 2 3 4]
	s5 := a1[4:]
	fmt.Println(s5) // [5 6]
	s6 := a1[:]
	fmt.Println(s6) // [1 2 3 4 5 6]
	// 切片容量指的是底层数组的容量
	// 数组容量是从起始切的位置一直到最后。
	fmt.Printf("len(s3):%d, cap(s3):%d\n", len(s3), cap(s3))
	// len(s3):4, cap(s3):6
	fmt.Printf("len(s5):%d, cap(s5):%d\n", len(s5), cap(s5))
	// len(s5):2, cap(s5):2

	// 切片是在数组上面的滑动窗口，引用类型
	// 对切片再切片也只是改变窗口的大小
	s7 := s5[1:]
	a1[5] = 100
	fmt.Println("s7", s7, "s5", s5) // s7 [100] s5 [5 100]

	// 3. make函数创建切片
	s8 := make([]int, 5, 10)
	fmt.Println(s8)      // [0 0 0 0 0]
	fmt.Println(len(s8)) //5
	fmt.Println(cap(s8)) //10

	s9 := []int{1, 2, 3}
	s10 := s9
	fmt.Println(s9, s10)
	s9[0] = 100
	fmt.Println(s9, s10)

	// append方法为切片添加元素
	var s11 []int
	for i := 0; i < 10; i++ {
		// 调用append必须使用之前的切片来接收返回值
		s11 = append(s11, i)
		fmt.Print("value：", s11, " len:", len(s11), " cap:", cap(s11))
		fmt.Printf(" ptr:%p\n", s11)
	}

	// copy
	s12 := []int{1,2,3,4,5}
	s13 := make([]int, 5, 5)
	copy(s13,s12)
	fmt.Println(s12)
	fmt.Println(s13)
	s13[0] = 100
	fmt.Println(s12)
	// [1 2 3 4 5]
	fmt.Println(s13)
	// [100 2 3 4 5]


	// 删除
	a2 := [...]int{1,2,3} // 数组，习惯数组的这个写法
	s14 := a2[:]		// 切片
	fmt.Println(s14)
	// 删除第2个
	s14 = append(s14[:1], s14[2:]...)
	fmt.Println(s14) // [1 3]
	fmt.Println(a2) // [1 3 3]
	

	// 
}
