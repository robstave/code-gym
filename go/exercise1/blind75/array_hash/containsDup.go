package utils

// note...the brute force is O(N2)
// and another way is to sort first then scan for doubles

// hasDuplicate tracks values in a map and returns as soon as it finds a repeat.
// Time: O(n) average; space: O(n), where n is the number of input values.
func hasDuplicate(nums []int) bool {

	valueMap := make(map[int]bool)

	for _, value := range nums {

		_, ok := valueMap[value]
		if !ok {
			valueMap[value] = true
		} else {
			return true
		}
	}

	return false
}

// hasDuplicate2 uses a set to detect repeats in one pass. The empty struct
// stores no per-entry value, unlike map[int]bool, and len(nums) preallocates
// capacity to reduce map growth for inputs without early duplicates.
// Time: O(n) average; space: O(n), where n is the number of input values.
func hasDuplicate2(nums []int) bool {
	seen := make(map[int]struct{}, len(nums))
	for _, value := range nums {
		if _, exists := seen[value]; exists {
			return true
		}
		seen[value] = struct{}{}
	}
	return false
}
