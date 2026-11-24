// Case 2: Applying a raise. Difficulty: Easy.
package main

import "fmt"

type employee struct {
	salary int
}

func applyRaise(employees []employee, amount int) {
	for _, employee := range employees {
		employee.salary += amount
	}
}

func main() {
	employees := []employee{{salary: 100}, {salary: 200}}
	applyRaise(employees, 10)
	fmt.Println(employees[0].salary, employees[1].salary)
}
