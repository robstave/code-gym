package main

import (
	"fmt"
	"time"
)

func main() {

	// create a channel

	ch := make(chan string, 2)

	go func() {

		time.Sleep(300 * time.Millisecond)

		ch <- "ping"
	}()

	go func() {

		time.Sleep(100 * time.Millisecond)

		ch <- "pong"
	}()

	// start a go proc and write to it

	val := <-ch

	fmt.Printf(" results: %s", val)

	val = <-ch

	fmt.Printf(" results: %s", val)

	//end

}
