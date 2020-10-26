package main

import (
	"fmt"
	"strings"
)

func main() {
	// 转意
	s1 := "\"你\\好'啊\""
	fmt.Println(s1)

	// 多行
	s2 := `hello 
		啊这
	`
	fmt.Println(s2)

	// 相关操作
	s3 := "/Users/qinzetao/workspace/go"
	s4 := strings.Split(s3, "/")
	fmt.Println(s4)
	fmt.Println(strings.Join(s4, "?"))

	fmt.Println("len", len(s3))

	fmt.Println(strings.Contains(s3, "Users"))

	fmt.Println(strings.HasPrefix(s3, "/U"))
	fmt.Println(strings.HasSuffix(s3, "go"))
	fmt.Println(strings.Index(s3, "e"))
	fmt.Println(strings.LastIndex(s3, "go"))


}
