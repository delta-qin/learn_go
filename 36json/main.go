package main

import (
	"encoding/json"
	"fmt"
)

type person struct {
	Name string `json:"name" db:"name" ini:"name"`
	Age  int    `json:"age"`
}

func main() {
	p1 := person{
		Name: "de",
		Age:  16,
	}

	b, err := json.Marshal(p1)
	if err != nil {
		fmt.Printf("err: %v", err)
		return
	}

	fmt.Printf("%v\n", string(b))

	// 字符串就是字节数组的切片
	str:= `{"name":"de", "age":12}`
	var p2 person
	json.Unmarshal([]byte(str), &p2)
	fmt.Printf("%#v\n", p2)
}
