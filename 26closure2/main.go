package main

import (
	"fmt"
	"strings"
)

func makeSuffix(suffix string) func(string) string {
	return func(name string) string {
		if !strings.HasSuffix(name, suffix) {
			return name + suffix
		}
		return name
	}
}

func main() {
	jpg := makeSuffix(".jpg")
	png := makeSuffix(".png")

	name1 := jpg("1")
	name2 := png("2")

	fmt.Println(name1)
	fmt.Println(name2)
}
