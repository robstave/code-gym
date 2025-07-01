// util.go
package util

// IsPalindrome reports whether s reads the same forward and backward.
func IsPalindrome(s string) bool {

	l := len(s)

	// size 0 is true
	if l <= 1 {
		return true
	}

	for i := 0; i < l/2-1; i++ {

		if s[i] != s[l-i-1] {
			return false
		}

	}
	return true
}

// IsPalindrome reports whether s reads the same forward and backward.
// the chinese chars mess it up...but this works
func IsPalindrome2(s string) bool {

	var r = []rune(s)
	l := len(r)

	// size 0 is true
	if l <= 1 {
		return true
	}

	for i := 0; i < l/2-1; i++ {

		if r[i] != r[l-i-1] {
			return false
		}

	}
	return true
}
