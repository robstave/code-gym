// group_anagrams.go
package util

import (
	"slices"
)

func Order(s string) string {

	runes := []rune(s)
	slices.Sort(runes)
	return string(runes)

}

// GroupAnagrams takes a slice of strings and returns a map where:
// - Each key is a canonical form of an anagram group
// - Each value is a slice of strings from the input that are anagrams of each other
//
// For example, given ["eat", "tea", "tan", "ate", "nat", "bat"],
// the function should return a map with three groups:
// - ["ate", "eat", "tea"]
// - ["nat", "tan"]
// - ["bat"]
//
// The order of strings within each group and the order of the groups themselves
// don't matter. You may decide how to represent the key in the map.
func GroupAnagrams(words []string) map[string][]string {

	result := make(map[string][]string)

	for _, v := range words {

		key := Order(v)

		result[key] = append(result[key], v)

	}

	return result
}

// ImproveSimilarityDetection takes GroupAnagrams a step further by identifying words
// that are "almost anagrams" - words that would be anagrams if one character was changed.
// For example, "listen" and "silent" are anagrams, but "listen" and "silenx" are almost anagrams.
//
// The function returns a map where:
// - Each key represents a group of similar words
// - Each value is a slice of strings that are either anagrams or "almost anagrams" of each other
//
// This requires more advanced string similarity detection.
func ImproveSimilarityDetection(words []string) map[string][]string {
	// Your implementation here
	return nil
}
