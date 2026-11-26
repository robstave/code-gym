package utils

func longestConsecutive(nums []int) int {
	values := make(map[int]struct{}, len(nums))
	for _, value := range nums {
		values[value] = struct{}{}
	}

	longest := 0
	for value := range values {
		if _, hasPrevious := values[value-1]; hasPrevious {
			continue
		}

		length := 1
		for next := value + 1; ; next++ {
			if _, exists := values[next]; !exists {
				break
			}
			length++
		}
		if length > longest {
			longest = length
		}
	}
	return longest
}

// passes but is n2
// bre
func longestConsecutiveNotRight(nums []int) int {

	myMap := make(map[int]struct{})

	for _, value := range nums {
		myMap[value] = struct{}{}
	}

	count := 0
	for _, value := range nums {

		if _, hasPrevious := myMap[value-1]; hasPrevious {
			continue
		}

		inner := 1

		for {
			value = value + 1
			_, ok := myMap[value]
			if ok {
				inner++
			} else {
				break
			}

		}
		if inner > count {
			count = inner
		}
	}
	return count
}
