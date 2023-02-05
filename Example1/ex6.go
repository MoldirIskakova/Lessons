package main

import (
	"fmt"
)

func main() {
	var a int
	fmt.Println("a manin engiz:")
	fmt.Scanf("%d\n", &a)

	var b int
	fmt.Println("b manin engiz:")
	fmt.Scanf("%d\n", &b)

	var c int
	fmt.Println("c manin engiz:")
	fmt.Scanf("%d\n", &c)

	V := a * b * c
	S := 2 * (a*b + b*c + a*c)
	fmt.Println("Paralipiped ob'em:", V)
	fmt.Println("Plochad' poverhnosti:", S)

}
