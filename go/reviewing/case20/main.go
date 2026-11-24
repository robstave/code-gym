// Case 20: Exporting a CSV row. Difficulty: Medium.
package main

import (
	"encoding/csv"
	"fmt"
	"strings"
)

func exportRow(fields []string) string {
	return strings.Join(fields, ",") + "\n"
}

func main() {
	row := exportRow([]string{"Ada, Jr.", "London"})
	fields, err := csv.NewReader(strings.NewReader(row)).Read()
	if err != nil {
		fmt.Println("Read error:", err)
		return
	}
	fmt.Println("Field count:", len(fields))
}
