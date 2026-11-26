package utils

import "testing"

func TestIsAnagram(t *testing.T) {
	tests := []struct {
		name string
		s    string
		t    string
		want bool
	}{
		{name: "anagrams", s: "racecar", t: "carrace", want: true},
		{name: "both strings empty", s: "", t: "", want: true},
		{name: "first string empty", s: "", t: "a", want: false},
		{name: "second string empty", s: "a", t: "", want: false},
		{name: "same length but mismatched characters", s: "ab", t: "ac", want: false},
		{name: "different lengths", s: "a", t: "ab", want: false},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := isAnagram(tt.s, tt.t); got != tt.want {
				t.Errorf("isAnagram(%q, %q) = %t, want %t", tt.s, tt.t, got, tt.want)
			}
			if got := isAnagram2(tt.s, tt.t); got != tt.want {
				t.Errorf("isAnagram2(%q, %q) = %t, want %t", tt.s, tt.t, got, tt.want)
			}
		})
	}
}
