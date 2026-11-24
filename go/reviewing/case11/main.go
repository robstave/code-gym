// Case 11: Product inventory. Difficulty: Easy.
package main

import "fmt"

func hasProduct(inventory map[string]int, product string) bool {

	return inventory[product] != 0
}

func main() {
	inventory := map[string]int{"pencils": 0}
	fmt.Println("Found:", hasProduct(inventory, "pencils"))
}
