package main

import "fmt"

func main() {
	a := 5
	b := 6
	c := 8

	V := a * b * c
	S := 2 * (a*b + b*c + a*c)
	fmt.Println("Paralipiped ob'em:", V)
	fmt.Println("Plochad' poverhnosti:", S)

}
