package main

import "fmt"

type employee struct {
	name   string
	sex    string
	age    int
	salary int
}

func newEmployee(name, sex string, age, salary int) employee {
	return employee{
		name:   name,
		sex:    sex,
		age:    age,
		salary: salary,
	}
}

func (e employee) getSalary(name string) int {
	if e.name == name {
		return e.salary
	} else {
		return 0
	}
}
func main() {
	employee1 := newEmployee("Vasya", "W", 40, 1000)
	employee2 := newEmployee("Zhan", "M", 35, 7000)

	//	fmt.Printf("%s\n", employee1.getinfo())
	//	fmt.Printf("%s\n", employee2.getinfo())
	fmt.Println(employee1.name, " get ", employee1.getSalary("Vasya"))

	fmt.Printf("%d\n", employee2.getSalary("Zhan"))
}
