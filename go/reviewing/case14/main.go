// Case 14: Retry delay. Difficulty: Easy.
package main

import (
	"fmt"
	"time"
)

func retryDelay(milliseconds int) time.Duration {
	return time.Duration(milliseconds)
}

func main() {
	fmt.Println("Delay:", retryDelay(500))
}
