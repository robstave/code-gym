// utils2.go
package util

// SumInts calculates the sum of all integers in a slice.
func SumInts(nums []int) int {

	var sum = 0

	for _, num := range nums {
		sum = sum + num
	}
	return sum
}

// Contains checks if a slice contains a specific element.
func Contains(slice []int, element int) bool {

	for _, val := range slice {
		if val == element {
			return true
		}
	}
	return false
}

// FilterEvenInts filters out all odd integers from a slice, leaving only even integers.
func FilterEvenInts(nums []int) []int {

	results := make([]int, 0, len(nums))

	for _, val := range nums {

		if val%2 == 0 {
			results = append(results, val)
		}
	}
	return results
}

// ReverseInts reverses the order of integers in a slice.
func ReverseInts(nums []int) []int {

	for i, j := 0, len(nums)-1; i < j; i, j = i+1, j-1 {
		nums[i], nums[j] = nums[j], nums[i]
	}
	return nums
}

func ReverseIntsWithAppend(nums []int) []int {
	results := make([]int, 0, len(nums))
	for i := len(nums) - 1; i >= 0; i-- {
		results = append(results, nums[i])
	}
	return results
}

// ReverseIntsRecursive reverses the order of integers in a slice using recursion.
func ReverseIntsRecursive(nums []int) []int {
	// Base case: empty slice or single element
	if len(nums) <= 1 {
		// Create a new slice to avoid modifying the original
		result := make([]int, len(nums))
		copy(result, nums)
		return result
	}

	// Recursive case: first element goes to the end, recurse on the rest
	return append(ReverseIntsRecursive(nums[1:]), nums[0])
}

// FindMax finds the maximum integer in a slice.
func FindMax(nums []int) int {
	if len(nums) == 0 {
		return 0
	}
	max := nums[0]
	for _, num := range nums {
		if num > max {
			max = num
		}
	}
	return max
}

// CountOccurrences counts the occurrences of a specific integer in a slice.
func CountOccurrences(nums []int, target int) int {
	count := 0
	for _, num := range nums {
		if num == target {
			count++
		}
	}
	return count
}

// MergeSlices merges two slices into one.
func MergeSlices(slice1, slice2 []int) []int {
	return append(slice1, slice2...)
}
