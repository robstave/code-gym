// Case 9: Event timestamp. Difficulty: Medium.
package main

import (
	"fmt"
	"time"
)

func eventTime(milliseconds int64) time.Time {
	return time.Unix(milliseconds, 0).UTC()
}

func main() {
	fmt.Println("Event:", eventTime(1700000000000).Format(time.RFC3339))
}
