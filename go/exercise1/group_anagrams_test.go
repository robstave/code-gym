// group_anagrams_test.go
package util

import (
	"reflect"
	"sort"
	"testing"
)

// Helper function to normalize the output for comparison
func normalizeAnagramOutput(result map[string][]string) [][]string {
	var normalized [][]string
	for _, group := range result {
		// Sort each group
		sort.Strings(group)
		normalized = append(normalized, group)
	}

	// Sort the groups by first element for consistent comparison
	sort.Slice(normalized, func(i, j int) bool {
		if len(normalized[i]) == 0 || len(normalized[j]) == 0 {
			return len(normalized[i]) < len(normalized[j])
		}
		return normalized[i][0] < normalized[j][0]
	})

	return normalized
}

func TestGroupAnagrams(t *testing.T) {
	tests := []struct {
		name  string
		input []string
		want  [][]string
	}{
		{
			name:  "basic example",
			input: []string{"eat", "tea", "tan", "ate", "nat", "bat"},
			want: [][]string{
				{"ate", "eat", "tea"},
				{"bat"},
				{"nat", "tan"},
			},
		},
		{
			name:  "case sensitivity",
			input: []string{"ABc", "abc", "BcA", "cab"},
			want: [][]string{
				{"ABc", "BcA"},
				{"abc", "cab"},
			},
		},
		{
			name:  "all anagrams",
			input: []string{"abc", "bac", "cab", "bca", "acb", "cba"},
			want: [][]string{
				{"abc", "acb", "bac", "bca", "cab", "cba"},
			},
		},

		{
			name:  "no anagrams",
			input: []string{"cat", "dog", "bird"},
			want: [][]string{
				{"bird"},
				{"cat"},
				{"dog"},
			},
		},
		/*
			{
				name:  "empty input",
				input: []string{},
				want:  [][]string{},
			},

				{
					name:  "special characters",
					input: []string{"a!bc", "ab!c", "bc!a", "cab!"},
					want: [][]string{
						{"a!bc"},
						{"ab!c"},
						{"bc!a"},
						{"cab!"},
					},
				},

					{
						name:  "varied lengths",
						input: []string{"a", "aa", "aaa", "aaaa", "a"},
						want: [][]string{
							{"a", "a"},
							{"aa"},
							{"aaa"},
							{"aaaa"},
						},
					},
		*/
	}

	for _, tc := range tests {
		t.Run(tc.name, func(t *testing.T) {
			got := GroupAnagrams(tc.input)
			gotNormalized := normalizeAnagramOutput(got)

			if !reflect.DeepEqual(gotNormalized, tc.want) {
				t.Errorf("GroupAnagrams() = %v, want %v", gotNormalized, tc.want)
			}
		})
	}
}

/*
func TestImproveSimilarityDetection(t *testing.T) {
	tests := []struct {
		name  string
		input []string
		want  [][]string
	}{
		{
			name:  "exact anagrams and near misses",
			input: []string{"listen", "silent", "silenx", "enlist", "inlest", "tinsel"},
			want: [][]string{
				{"enlist", "inlest", "listen", "silent", "tinsel"},
				{"silenx"},
			},
		},
		{
			name:  "single letter differences",
			input: []string{"hello", "hallo", "holla", "olleh", "ollah"},
			want: [][]string{
				{"hello", "olleh"},
				{"hallo", "holla", "ollah"},
			},
		},
		{
			name:  "special case - mixed difficulty",
			input: []string{"word", "wrdo", "wodr", "sword", "drowning"},
			want: [][]string{
				{"word", "wodr", "wrdo"},
				{"sword"},
				{"drowning"},
			},
		},
		{
			name:  "empty input",
			input: []string{},
			want:  [][]string{},
		},
		{
			name:  "no similarity",
			input: []string{"cat", "dog", "fish", "bird"},
			want: [][]string{
				{"bird"},
				{"cat"},
				{"dog"},
				{"fish"},
			},
		},
	}

	for _, tc := range tests {
		t.Run(tc.name, func(t *testing.T) {
			got := ImproveSimilarityDetection(tc.input)
			gotNormalized := normalizeAnagramOutput(got)

			// Check if each group in the expected result is present in the actual result
			if !reflect.DeepEqual(gotNormalized, tc.want) {
				t.Errorf("ImproveSimilarityDetection() = %v, want %v", gotNormalized, tc.want)
			}
		})
	}
}
*/
