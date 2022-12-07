package main

import "fmt"

func main() {
	employee := struct {
		name   string
		sex    string
		age    int
		salary int
	}{
		name:   "Moldir",
		sex:    "w",
		age:    25,
		salary: 1000000,
	}
	fmt.Printf("%+v\n", employee)
}
