package main

import (
	"fmt"
)

func main() {

	nums := []int{3, 6, 7, 3, 4, 1, 3}
	mp := make(map[int]int)
	for _, n := range nums {
		mp[n]++
	}
	fmt.Println(mp)

}
