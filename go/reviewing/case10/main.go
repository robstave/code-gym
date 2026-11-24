// Case 10: Removing canceled orders. Difficulty: Medium.
package main

import "fmt"

func removeCanceled(orders []string) []string {
	for i := 0; i < len(orders); i++ {
		if orders[i] == "canceled" {
			orders = append(orders[:i], orders[i+1:]...)
		}
	}
	return orders
}

func main() {
	orders := []string{"active", "canceled", "canceled", "active"}
	fmt.Println("Remaining:", removeCanceled(orders))
}
