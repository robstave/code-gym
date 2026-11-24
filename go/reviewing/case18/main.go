// Case 18: Importing lines. Difficulty: Medium.
package main

import (
	"bufio"
	"errors"
	"fmt"
	"io"
	"strings"
)

type failingReader struct {
	first io.Reader
}

func (r *failingReader) Read(p []byte) (int, error) {
	n, err := r.first.Read(p)
	if err == io.EOF {
		return 0, errors.New("input failed")
	}
	return n, err
}

func countLines(input io.Reader) (int, error) {
	scanner := bufio.NewScanner(input)
	count := 0
	for scanner.Scan() {
		count++
	}
	return count, nil
}

func main() {
	input := &failingReader{first: strings.NewReader("one\n")}
	count, err := countLines(input)
	fmt.Println("Lines:", count, "Error:", err)
}
