package main

import "fmt"

type sotrudniki struct {
	name     string
	age      int
	salary   int
	position string
}

func createTable(name, position string, age, salary int) sotrudniki {
	return sotrudniki{
		name:     name,
		age:      age,
		salary:   salary,
		position: position,
	}
}

func (r sotrudniki) getAge(name string) int {
	if r.name == name {
		return r.age
	} else {
		return 0
	}
}

func (zp sotrudniki) getSalary(name string) int {
	if zp.name == name {
		return zp.salary
	} else {
		return 0
	}

}

func main() {

	sotrudnik1 := createTable("Lola", "HR", 50, 80000)
	sotrudnik2 := createTable("Max", "Master", 40, 50000)

	fmt.Println("Age Lola is", sotrudnik1.getAge("Lola"), "zp", sotrudnik1.getSalary("Lola"))
	fmt.Println("Age Max is", sotrudnik2.getAge("Max"), "zp", sotrudnik2.getSalary("Max"))
}
