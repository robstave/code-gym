package util

import "testing"

func TestReverse(t *testing.T) {
	tests := []struct {
		input string
		want  string
	}{
		{"hello", "olleh"},
		{"", ""},
		{"a", "a"},
		{"abc", "cba"},
		{"racecar", "racecar"},
	}
	for _, tc := range tests {
		got := Reverse1(tc.input)
		if got != tc.want {
			t.Errorf("Reverse(%q) = %q; want %q", tc.input, got, tc.want)
		}
	}
}

func TestReverse2(t *testing.T) {
	tests := []struct {
		input string
		want  string
	}{
		{"hello", "olleh"},
		{"", ""},
		{"a", "a"},
		{"abc", "cba"},
		{"racecar", "racecar"},
	}
	for _, tc := range tests {
		got := Reverse2(tc.input)
		if got != tc.want {
			t.Errorf("Reverse(%q) = %q; want %q", tc.input, got, tc.want)
		}
	}
}
