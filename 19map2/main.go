package main

import (
	"fmt"
	"math/rand"
	"sort"
)

func main() {
	scores := make(map[string]int, 200)

	for i := 0; i < 100; i++ {
		key := fmt.Sprintf("stu%02d", i)
		value := rand.Intn(100)
		scores[key] = value
	}

	fmt.Println(scores)

	s := make([]string, 0, 200)
	for k := range scores {
		s = append(s, k)
	}

	fmt.Println(s)

	sort.Strings(s)

	for _, key := range s {
		fmt.Println(key, scores[key])
	}

}
