// Case 12: JSON response. Difficulty: Easy.
package main

import (
	"encoding/json"
	"fmt"
)

type person struct {
	Name string `json:"name"`
	Age  int    `json:"age"`
}

func main() {
	data, err := json.Marshal(person{Name: "Ada", Age: 37})
	if err != nil {
		fmt.Println("Encoding error:", err)
		return
	}
	fmt.Println(string(data))
}
