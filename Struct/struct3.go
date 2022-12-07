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

func (e employee) getinfo() string {
	return fmt.Sprintf("Сотрудник:%s\nВозраст: %d\nЗарплата: %d\n", e.name, e.age, e.salary)
}
func main() {
	employee1 := newEmployee("Vasya", "W", 40, 1000)
	employee2 := newEmployee("Zhan", "M", 35, 7000)

	fmt.Printf("%s\n", employee1.getinfo())
	fmt.Printf("%s\n", employee2.getinfo())
}
