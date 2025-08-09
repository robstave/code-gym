// matrix_rotation.go
package util

// RotateMatrix90 takes an n x n matrix and rotates it 90 degrees clockwise.
// For example, given:
// [
//   [1, 2, 3],
//   [4, 5, 6],
//   [7, 8, 9]
// ]
// It should return:
// [
//   [7, 4, 1],
//   [8, 5, 2],
//   [9, 6, 3]
// ]
// Assume the matrix is square (n x n).
func RotateMatrix90(matrix [][]int) [][]int {
	// Your implementation here

	// check for [][]  in this case len =0
	if len(matrix) == 0 {
		return matrix
	}
	n := len(matrix[0])

	if n <= 1 {
		return matrix
	}

	// note you have to initialize the whole thing
	result := make([][]int, n)

	for i := range result {
		result[i] = make([]int, n)
	}

	for i := 0; i <= n-1; i++ {

		for j := 0; j <= n-1; j++ {

			result[j][n-i-1] = matrix[i][j]

		}
	}
	return result
}

// RotateMatrixInPlace90 is similar to RotateMatrix90 but modifies the matrix in-place
// without allocating a new matrix.
// This is more challenging as you need to rotate elements without using extra space.
func RotateMatrixInPlace90(matrix [][]int) {
	// Your implementation here

	// check for [][]  in this case len =0
	if len(matrix) == 0 {
		return
	}
	n := len(matrix[0])

	if n <= 1 {
		return
	}

	for i := 0; i <= n-1; i++ {

		a := matrix[i][i]
		b := matrix[i][n-i-1]
		c := matrix[n-i-1][i]
		d := matrix[n-i-1][n-i-1]

		matrix[i][i] = b
		matrix[i][n-i-1] = c
		matrix[n-i-1][i] = d
		matrix[n-i-1][n-i-1] = a

	}

}
