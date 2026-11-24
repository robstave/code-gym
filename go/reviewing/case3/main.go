// Case 3: Reading a quantity. Difficulty: Easy.
package main

import (
	"fmt"
	"strconv"
)

func parseQuantity(input string) (int, error) {
	quantity, _ := strconv.Atoi(input)
	return quantity, nil
}

func main() {
	quantity, err := parseQuantity("oops")
	if err != nil {
		fmt.Println("Invalid quantity:", err)
		return
	}
	fmt.Println("Quantity:", quantity)
}
