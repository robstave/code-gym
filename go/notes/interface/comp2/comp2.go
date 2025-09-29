package main

import "fmt"

type names struct {
	first string
	last  string
}

type stats struct {
	sex string
	age int
}

type person struct {
	name names
	stat stats
}

func newPerson(first string, last string, sex string, age int) *person {

	p := person{
		name: names{
			first: first,
			last:  last,
		},
		stat: stats{
			sex: sex,
			age: age,
		},
	}

	return &p
}

// note you do not have to use different names for the embedded structs
type person2 struct {
	names
	stats
}

func newPerson2(first string, last string, sex string, age int) *person2 {

	p := person2{
		names: names{
			first: first,
			last:  last,
		},
		stats: stats{
			sex: sex,
			age: age,
		},
	}

	return &p
}

func main() {

	jd := newPerson("Jon", "Doe", "M", 33)

	fmt.Println(jd)
	fmt.Println("First name:", jd.name.first)

	jd2 := newPerson2("Jon2", "Doe2", "F", 11)

	fmt.Println(jd2)
	fmt.Println("First name:", jd2.names.first)

}
