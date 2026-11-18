package util

import (
	"fmt"
	"reflect"
	"testing"
)

func Example_combine() {
	combined := combine(
		func(value int) { fmt.Println("first:", value) },
		func(value int) { fmt.Println("second:", value) },
	)
	combined(3)

	// Output:
	// first: 3
	// second: 3
}

func Example_combineGenerics() {
	combined := combineGenerics(
		func(value int) { fmt.Println("number:", value) },
		func(value string) { fmt.Println("text:", value) },
	)
	combined(42, "hello")

	// Output:
	// number: 42
	// text: hello
}

func TestCombine(t *testing.T) {
	var calls []int
	combined := combine(
		func(value int) { calls = append(calls, value) },
		func(value int) { calls = append(calls, value*2) },
	)
	if len(calls) != 0 {
		t.Fatal("combine called its functions before the returned function was invoked")
	}
	combined(3)
	combined(4)

	want := []int{3, 6, 4, 8}
	if !reflect.DeepEqual(calls, want) {
		t.Errorf("calls = %v, want %v", calls, want)
	}
}

func TestCombineGenerics(t *testing.T) {
	var calls []string
	f1 := func(value int) {
		if value != 42 {
			t.Errorf("f1 received %d, want 42", value)
		}
		calls = append(calls, "f1")
	}
	f2 := func(value string) {
		if value != "hello" {
			t.Errorf("f2 received %q, want %q", value, "hello")
		}
		calls = append(calls, "f2")
	}

	combined := combineGenerics(f1, f2)
	if len(calls) != 0 {
		t.Fatal("combineGenerics called its functions before the returned function was invoked")
	}
	combined(42, "hello")
	combined(42, "hello")

	want := []string{"f1", "f2", "f1", "f2"}
	if !reflect.DeepEqual(calls, want) {
		t.Errorf("calls = %v, want %v", calls, want)
	}
}

func TestCombineGenericsSameType(t *testing.T) {
	var calls []int
	combined := combineGenerics(
		func(value int) { calls = append(calls, value) },
		func(value int) { calls = append(calls, value) },
	)
	combined(1, 2)

	want := []int{1, 2}
	if !reflect.DeepEqual(calls, want) {
		t.Errorf("calls = %v, want %v", calls, want)
	}
}

func TestFunctions(t *testing.T) {
	tests := []struct {
		input [][]int
		want  []int
	}{
		{[][]int{{1, 2}, {3, 4}}, []int{1, 2, 3, 4}},
		{[][]int{}, []int{}},
		{[][]int{{}, {5, 6}, {}}, []int{5, 6}},
		{[][]int{{7}, {8, 9}, {10, 11, 12}}, []int{7, 8, 9, 10, 11, 12}},
	}
	for _, tc := range tests {
		got := FlattenInts2(tc.input)
		if !equalSlices(got, tc.want) {
			t.Errorf("FlattenInts(%v) = %v; want %v", tc.input, got, tc.want)
		}
	}
}
