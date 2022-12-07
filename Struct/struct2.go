package main

import "fmt"

func main() {
	type employee struct {
		name   string
		sex    string
		age    int
		salary int
	}

	employee1 := employee{
		name:   "Vasya",
		sex:    "W",
		age:    40,
		salary: 500,
	}
	employee2 := employee{
		name:   "Zhan",
		sex:    "M",
		age:    35,
		salary: 1000,
	}
	fmt.Printf("%+v\n", employee1)
	fmt.Printf("%+v\n", employee2)
}
