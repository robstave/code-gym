package utils

import "math/bits"

// hammingWeight counts the set bits in a non-negative integer by examining
// each bit from least significant to most significant.
// this is not exactly the cleanest, but works
// time  O(logn)
// space O(1)
func hammingWeight(value int) int {

	var count int = 0
	for value != 0 {

		out := value & 1
		if out == 1 {
			count++
		}
		value = value >> 1
	}
	return count

}

// hammingWeightIdiomatic counts set bits by repeatedly clearing the lowest set bit.
// Negative values are counted using their machine-width integer representation.
func hammingWeightIdiomatic(value int) int {
	count := 0
	for value != 0 {
		value &= value - 1
		count++
	}
	return count
}

// hammingWeightBit counts the set bits using Go's standard library.
// Negative values are interpreted using their machine-width unsigned representation.
func hammingWeightBit(value int) int {
	return bits.OnesCount(uint(value))
}

// Other suggestions:
//   - Use uint as the parameter type, or otherwise define negative-value behavior consistently.
//   - Keep both implementations only if comparing the manual and standard-library approaches is useful.
//   - Add test cases for negative values once their expected behavior is defined.
