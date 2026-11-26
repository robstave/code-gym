package utils

import (
	"sort"
	"strings"
)

func isAnagram(s string, t string) bool {

	sSlice := strings.Split(s, "")
	tSlice := strings.Split(t, "")

	sort.Strings(sSlice)
	sort.Strings(tSlice)

	ss := strings.Join(sSlice, "")
	ts := strings.Join(tSlice, "")
	return ss == ts

}

func isAnagram2(s string, t string) bool {

	ls := len(s)
	lt := len(t)
	if ls != lt {
		return false
	}

	maps := make(map[rune]int)

	for _, rune_s := range s {

		value, ok := maps[rune_s]
		if ok {
			maps[rune_s] = value + 1
		} else {
			maps[rune_s] = 0
		}

	}

	for _, rune_t := range t {

		value, ok := maps[rune_t]
		if ok {
			maps[rune_t] = value - 1
		} else {
			return false
		}
		if value < 0 {
			return false
		}

	}

	return true

}
