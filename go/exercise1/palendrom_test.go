// util_test.go
package util

import "testing"

func TestIsPalindrome(t *testing.T) {
	tests := []struct {
		input string
		want  bool
	}{
		{"racecar", true},
		{"hello", false},
		{"", true},
		{"a", true},
		{"1221", true},
		//{"世界界世", true},
		{"GoLang", false},
	}
	for _, tc := range tests {
		got := IsPalindrome(tc.input)
		if got != tc.want {
			t.Errorf("IsPalindrome(%q) = %v; want %v", tc.input, got, tc.want)
		}
	}
}

func TestIsPalindrome2(t *testing.T) {
	tests := []struct {
		input string
		want  bool
	}{
		{"racecar", true},
		{"hello", false},
		{"", true},
		{"a", true},
		{"1221", true},
		{"世界界世", true},
		{"GoLang", false},
	}
	for _, tc := range tests {
		got := IsPalindrome2(tc.input)
		if got != tc.want {
			t.Errorf("IsPalindrome(%q) = %v; want %v", tc.input, got, tc.want)
		}
	}
}
