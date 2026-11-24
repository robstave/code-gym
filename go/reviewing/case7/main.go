// Case 7: Request timeout. Difficulty: Medium.
package main

import (
	"context"
	"fmt"
	"time"
)

func requestContext(parent context.Context) context.Context {
	ctx, cancel := context.WithTimeout(parent, time.Second)
	defer cancel()
	return ctx
}

func main() {
	ctx := requestContext(context.Background())
	fmt.Println("Context error:", ctx.Err())
}
