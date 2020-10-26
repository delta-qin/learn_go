package main

import "fmt"

func main() {
	s := "deltaqin你好"

	fmt.Println(len(s))

	for i := 0; i < len(s); i++ {
		fmt.Printf("%v(%c)", s[i], s[i])
	}
	fmt.Println()

	for _, r := range s {
		fmt.Printf("%v(%c)", r, r)
	}

	fmt.Println("================")
	// 100(d)101(e)108(l)116(t)97(a)113(q)105(i)110(n)228(ä)189(½)160( )229(å)165(¥)189(½)
	// 100(d)101(e)108(l)116(t)97(a)113(q)105(i)110(n)20320(你)22909(好)

	byteS1 := []byte(s)
	byteS1[0]='a'
	fmt.Println(string(byteS1))

	byteS2 := []rune(s)
	// 注意字符不可以使用""
	byteS2[0] = '秦'
	fmt.Println(string(byteS2))
}
