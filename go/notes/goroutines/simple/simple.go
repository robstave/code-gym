package main

import (
	"fmt"
	"time"
)

func f(s string) {
	for i := range 5 {
		fmt.Printf(" %s: %d\n", s, i)
	}
}

func main() {

	f("direct")

	go f("goroutine")

	go func(msg string) {
		fmt.Println(msg)
	}("going")

	time.Sleep(time.Second)
	fmt.Println("done")
}
