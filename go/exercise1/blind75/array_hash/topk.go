package utils

import (
	"sort"
)

func topKFrequent(nums []int, k int) []int {

	myMap := make(map[int]int)

	// map out freq.  you dont have to check for ok
	for _, num := range nums {
		myMap[num]++
	}
	//log.Printf("myMap: %v", myMap)

	arr := make([][2]int, 0)
	for key, count := range myMap {
		arr = append(arr, [2]int{count, key})
	}
	//log.Printf("arr: %v", arr)
	// Each entry is [count, value]. sort.Slice places i before j when the
	// comparison is true, so greater counts come first. Equal counts have no
	// guaranteed relative order.
	sort.Slice(arr, func(i, j int) bool {
		return arr[i][0] > arr[j][0]
	})
	///.Printf("arr d: %v", arr)
	result := make([]int, 0)

	for i := range k {
		result = append(result, arr[i][1])
	}
	return result
}
