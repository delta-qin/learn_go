package main

import "fmt"

func main() {

	fmt.Println("========1 初始化不分配内存得到nil=======")
	var s map[string]string
	fmt.Println(s == nil)
	s = make(map[string]string, 10)
	s["name"] = "deltaqin"
	fmt.Println(s)

	fmt.Println("========2 直接分配内存获取map=======")
	scores := make(map[string]int, 8)
	scores["de"] = 90
	scores["lta"] = 100
	fmt.Println(scores)
	fmt.Println(scores["de"])
	fmt.Printf("%T\n", scores)

	fmt.Println("========3 新建map时候直接赋值=======")
	s1 := map[string]string{
		"name": "deltqin",
		"pass": "123",
	}
	fmt.Println(s1)

	fmt.Println("========4 打印map里面的值=======")
	s2 := make(map[string]string, 10)
	val, ok := s2["de"]
	if !ok {
		fmt.Println("查无此key")
	} else {
		fmt.Println(val)
	}
	fmt.Println(s2)

	fmt.Println("========5 遍历=======")
	b1 := make(map[string]string, 10)
	b1["a"] = "1"
	b1["b"] = "2"
	// 只k
	for k := range b1 {
		fmt.Println(k)
	}
	// 只v
	for _, v := range b1 {
		fmt.Println(v)
	}

	delete(b1, "a")
	fmt.Println(b1)

}
