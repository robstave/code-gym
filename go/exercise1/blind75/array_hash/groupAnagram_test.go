package utils

import (
	"slices"
	"testing"
)

func TestGroupAnagrams(t *testing.T) {
	tests := []struct {
		name string
		strs []string
		want [][]string
	}{
		{
			name: "example",
			strs: []string{"act", "pots", "tops", "cat", "stop", "hat"},
			want: [][]string{{"hat"}, {"act", "cat"}, {"stop", "pots", "tops"}},
		},
		{name: "empty input", strs: []string{}, want: [][]string{}},
		{name: "empty strings", strs: []string{"", "", "a"}, want: [][]string{{"", ""}, {"a"}}},
		{name: "no anagrams", strs: []string{"cat", "dog", "bird"}, want: [][]string{{"cat"}, {"dog"}, {"bird"}}},
		{name: "all anagrams", strs: []string{"abc", "cab", "bca"}, want: [][]string{{"abc", "cab", "bca"}}},
		{name: "duplicate words", strs: []string{"eat", "tea", "eat"}, want: [][]string{{"eat", "tea", "eat"}}},
		{name: "different character counts", strs: []string{"a", "aa", "ab", "aab", "baa"}, want: [][]string{{"a"}, {"aa"}, {"ab"}, {"aab", "baa"}}},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := groupAnagrams(tt.strs)
			for _, group := range got {
				slices.Sort(group)
			}
			for _, group := range tt.want {
				slices.Sort(group)
			}
			slices.SortFunc(got, slices.Compare[[]string])
			slices.SortFunc(tt.want, slices.Compare[[]string])
			if !slices.EqualFunc(got, tt.want, slices.Equal[[]string]) {
				t.Errorf("groupAnagrams(%q) = %q, want %q", tt.strs, got, tt.want)
			}
		})
	}
}
