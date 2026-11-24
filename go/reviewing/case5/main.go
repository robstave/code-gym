// Case 5: Counter. Difficulty: Easy.
package main

import "fmt"

type counter struct {
	value int
}

func (c counter) increment() {
	c.value++
}

func main() {
	c := counter{}
	c.increment()
	fmt.Println("Count:", c.value)
}
