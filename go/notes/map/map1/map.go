package main

import (
	"fmt"
	"sort"
	"strings"
)

/*
MAPS EXERCISE PLAYGROUND
========================

This file walks through core and advanced map usage in Go.
Sections contain demonstrations and a set of TODOs at the bottom to implement.

Covered concepts:
  1. Zero value vs make
  2. Insert / update / lookup / delete / existence check
  3. Iteration order (randomized) & extracting keys deterministically
  4. Maps of structs vs maps of pointers
  5. Counting / frequency maps
  6. Nested maps (map[string]map[string]int)
  7. Using map as a set
  8. Immutability patterns (copy-on-write, defensive copy)
  9. Value vs reference semantics (remember: map is reference type)
 10. Merging maps
 11. Filtering maps
 12. Inverting a map (value -> list of keys)
 13. Safe get with default
 14. Case-insensitive key handling layer

TODO EXERCISES appear near the end with stubs you fill in.
Run with: go run . (from folder containing this file as main module)
*/

// 1. Zero value vs make
func zeroVsMake() {
	var mNil map[string]int        // nil map: cannot assign (panic); len=0
	mMade := make(map[string]int)  // ready to use
	fmt.Println("mNil == nil?", mNil == nil, "len=", len(mNil))
	fmt.Println("mMade == nil?", mMade == nil, "len=", len(mMade))
	// Safe read from nil map returns zero value
	fmt.Println("read from nil map: mNil[foo] ->", mNil["foo"]) // 0
	// Writing to nil map would panic: uncomment to observe
	// mNil["x"] = 1
}

// 2. Insert / update / lookup / delete / existence
func basicOps() {
	m := make(map[string]int)
	m["apple"] = 3       // insert
	m["banana"] = 5
	m["apple"] = 4       // update
	v1 := m["apple"]     // lookup
	vMissing := m["pear"] // missing returns zero (0)
	_, ok := m["banana"] // existence check (only need ok)
	delete(m, "banana")
	_, okAfter := m["banana"]
	fmt.Println("apple=", v1, "pear=", vMissing, "banana existed?", ok, "after delete?", okAfter)
}

// 3. Iteration order is randomized; gather keys & sort for deterministic output
func orderedPrint(m map[string]int) {
	keys := make([]string, 0, len(m))
	for k := range m { keys = append(keys, k) }
	sort.Strings(keys)
	for _, k := range keys { fmt.Printf("%s => %d\n", k, m[k]) }
}

// 4. Map of structs vs map of pointers
type User struct { ID int; Name string }

func structVsPointerMap() {
	byIDVal := map[int]User{1:{1,"Ann"},2:{2,"Bob"}}
	byIDPtr := map[int]*User{3:{3,"Cat"},4:{4,"Don"}}
	// Updating field on value requires re-assignment
	u := byIDVal[1]; u.Name = "ANN"; byIDVal[1] = u
	// Pointer map allows direct mutation
	byIDPtr[3].Name = "CAT"
	fmt.Println("value map: ", byIDVal)
	fmt.Println("pointer map:", byIDPtr)
}

// 5. Frequency counting
func frequency(words []string) map[string]int {
	counts := make(map[string]int, len(words)) // capacity hint
	for _, w := range words { counts[w]++ }
	return counts
}

// 6. Nested maps
func nestedMapDemo() {
	scores := map[string]map[string]int{}
	ensureInner := func(category string) map[string]int {
		inner, ok := scores[category]
		if !ok { inner = make(map[string]int); scores[category] = inner }
		return inner
	}
	ensureInner("math")["alice"] += 10
	ensureInner("math")["bob"] += 7
	ensureInner("eng")["alice"] += 8
	fmt.Println("nested:", scores)
}

// 7. Map as a set (empty struct uses zero bytes)
type void = struct{}
var present void

func setDemo() {
	set := make(map[string]void)
	add := func(k string){ set[k]=present }
	contains := func(k string) bool { _, ok := set[k]; return ok }
	add("apple"); add("banana")
	fmt.Println("contains apple?", contains("apple"), "contains pear?", contains("pear"))
}

// 8. Immutability pattern (defensive copy before exposing)
func copyMap[K comparable, V any](m map[K]V) map[K]V {
	if m == nil { return nil }
	out := make(map[K]V, len(m))
	for k,v := range m { out[k]=v }
	return out
}

// 9. Reference semantics demonstration
func referenceSemantics() {
	m := map[string]int{"a":1}
	alias := m
	alias["a"] = 42
	fmt.Println("m[\"a\"] after modifying alias:", m["a"]) // 42
}

// 10. Merge maps (second overrides first)
func merge[K comparable, V any](a, b map[K]V) map[K]V {
	out := make(map[K]V, len(a)+len(b))
	for k,v := range a { out[k]=v }
	for k,v := range b { out[k]=v }
	return out
}

// 11. Filter map
func filterMap[K comparable, V any](m map[K]V, keep func(K,V) bool) map[K]V {
	out := make(map[K]V)
	for k,v := range m { if keep(k,v) { out[k]=v } }
	return out
}

// 12. Invert map[string]string -> map[string][]string (group keys by value)
func invertMulti(m map[string]string) map[string][]string {
	inv := make(map[string][]string)
	for k,v := range m { inv[v] = append(inv[v], k) }
	return inv
}

// 13. Safe get with default
func getOr[K comparable, V any](m map[K]V, k K, def V) V {
	if v, ok := m[k]; ok { return v }
	return def
}

// 14. Case-insensitive wrapper: stores keys in lower-case internally
type CIMap struct { data map[string]string }

func NewCIMap() *CIMap { return &CIMap{data: make(map[string]string)} }
func (c *CIMap) Set(k, v string) { c.data[strings.ToLower(k)] = v }
func (c *CIMap) Get(k string) (string, bool) { v, ok := c.data[strings.ToLower(k)]; return v, ok }
func (c *CIMap) Keys() []string { ks := make([]string,0,len(c.data)); for k := range c.data { ks = append(ks,k)}; sort.Strings(ks); return ks }

// ------------------------- TODO EXERCISES -------------------------
// Implement these for practice. Uncomment calls in main to test.

// TODO 1: topN(words []string, n int) []string
// Return the top n words by descending frequency (stable lexicographic tie-breaker).
// Use frequency map + slice sort. If n > unique words, return all.
func topN(words []string, n int) []string { panic("implement me") }

// TODO 2: intersectKeys(a,b map[string]int) []string
// Return sorted slice of keys present in both maps.
func intersectKeys(a, b map[string]int) []string { panic("implement me") }

// TODO 3: diffCounts(a,b map[string]int) map[string]int
// For every key in either map, result[key] = a[key] - b[key] (missing -> 0)
func diffCounts(a, b map[string]int) map[string]int { panic("implement me") }

// TODO 4: groupByLength(words []string) map[int][]string
// Keys are word lengths; values are the words of that length (preserve input order inside each bucket).
func groupByLength(words []string) map[int][]string { panic("implement me") }

// TODO 5: reverseIndex(docs []string) map[string][]int
// Build an inverted index mapping lower-cased tokens to document indices where they appear.
// Split on spaces, ignore empty tokens, do not de-duplicate (list may contain repeats for multiple occurrences).
func reverseIndex(docs []string) map[string][]int { panic("implement me") }

// ------------------------- MAIN DEMO -------------------------
func main() {
	fmt.Println("== zero vs make =="); zeroVsMake()
	fmt.Println("\n== basic ops =="); basicOps()
	fmt.Println("\n== ordered iteration ==")
	orderedPrint(map[string]int{"b":2,"a":1,"c":3})
	fmt.Println("\n== struct vs pointer map =="); structVsPointerMap()
	fmt.Println("\n== frequency =="); fmt.Println(frequency([]string{"go","is","fun","go","go","Is"}))
	fmt.Println("\n== nested map =="); nestedMapDemo()
	fmt.Println("\n== set demo =="); setDemo()
	fmt.Println("\n== reference semantics =="); referenceSemantics()
	fmt.Println("\n== merge =="); fmt.Println(merge(map[string]int{"a":1}, map[string]int{"a":2,"b":3}))
	fmt.Println("\n== filter =="); fmt.Println(filterMap(map[string]int{"a":1,"b":2,"c":3}, func(k string,v int) bool { return v%2==1 }))
	fmt.Println("\n== invert multi =="); fmt.Println(invertMulti(map[string]string{"u1":"teamA","u2":"teamB","u3":"teamA"}))
	fmt.Println("\n== getOr =="); fmt.Println(getOr(map[string]int{"x":10}, "y", 99))
	fmt.Println("\n== case-insensitive map ==")
	ci := NewCIMap(); ci.Set("Host","LOCAL"); ci.Set("host","override"); val, exists := ci.Get("HOST"); fmt.Println(ci.Keys(), val, exists)

	// Uncomment after implementing TODOs:
	// fmt.Println("\n== topN =="); fmt.Println(topN([]string{"a","b","a","c","b","a"}, 2))
	// fmt.Println("\n== intersectKeys =="); fmt.Println(intersectKeys(map[string]int{"a":1,"b":2}, map[string]int{"b":3,"c":4}))
	// fmt.Println("\n== diffCounts =="); fmt.Println(diffCounts(map[string]int{"a":5,"b":2}, map[string]int{"b":3,"d":1}))
	// fmt.Println("\n== groupByLength =="); fmt.Println(groupByLength([]string{"go","maps","are","cool"}))
	// fmt.Println("\n== reverseIndex =="); fmt.Println(reverseIndex([]string{"Go is cool","Maps are cool"}))
}

