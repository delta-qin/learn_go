package main

import (
	"fmt"
	"strings"
	"unicode"
)

func main() {
	// 1 计算汉字出现次数 ===========================
	s1 := "deltaqin你好哈哈哈"

	var count int
	for _, v := range s1 {
		if unicode.Is(unicode.Han, v) {
			count++
		}
	}
	fmt.Println(count)

	// 2 统计词频 ==================================
	s2 := "what do you want to do"
	s3 := strings.Split(s2, " ")
	fmt.Println(s2) // what do you want to do
	fmt.Println(s3) // [what do you want to do]
	m1 := make(map[string]int, 10)
	for _, w := range s3 {
		// fmt.Println(w)
		if _, ok := m1[w]; !ok {
			m1[w] = 1
		} else {
			m1[w]++
		}
	}
	// fmt.Println(m1)
	for k, v := range m1 {
		fmt.Println(k, v)
	}

	// 3 回文判断 =====================================
	// 上海自来水来自海上
	ss1 := "上海自来水来自海上"
	fmt.Println(len(ss1))
	// 27 = 9 * 3，一个字符3个字节
	// len直接获取到的是字符串对应的字节数
	// 转换为rune
	r := make([]rune, 0, len(ss1))
	for _, v := range ss1 {
		r = append(r, v)
	}
	fmt.Println(r) // [19978 28023 33258 26469 27700 26469 33258 28023 19978]

	rlen := len(r)
	for i := 0; i < rlen/2; i++ {
		// 直接使用首尾比较即可
		if r[i] != r[rlen-i-1]{
			fmt.Println("不是回文")
			return
		}
	}
	fmt.Println("是回文")


}
