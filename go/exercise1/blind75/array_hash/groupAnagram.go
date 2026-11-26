package utils

import (
	"slices"
	"strings"
)

func groupAnagrams(strs []string) [][]string {

	// create a map of sorted keys
	// add the value to each
	//spit them out

	anagramMap := make(map[string][]string)

	for _, value := range strs {

		// create key
		word := strings.Split(value, "")
		slices.Sort(word)
		key := strings.Join(word, "")

		// if key exists, append, else add
		_, ok := anagramMap[key]
		if ok {
			anagramMap[key] = append(anagramMap[key], value)

		} else {
			anagramMap[key] = []string{value}
		}
	}

	results := make([][]string, 0)

	for _, value := range anagramMap {
		results = append(results, value)

	}

	return results
}
