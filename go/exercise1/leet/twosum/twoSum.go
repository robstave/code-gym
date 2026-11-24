// You can edit this code!
// Click here and start typing.
package utils

import (
	"slices"
)

func twoSum(arr []int, target int) []int {

	//arr[x] + arr[y] = total
	// arr[y] = total - arr[x]
	// so map a map of arr[y] to index
	// then loop through (target - arr[x]) and see if it exists in map
	// that gives us x and y

	xMap := make(map[int]int)

	for i, value := range arr {
		xMap[value] = i
	}

	for j, v := range arr {
		comp := target - v

		index, ok := xMap[comp]

		if ok && (j != index) {
			result := []int{j, index}
			slices.Sort(result)
			return result
		}

	}

	return nil

}
