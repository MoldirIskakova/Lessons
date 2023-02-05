package main

import "fmt"

func main() {
	var a int
	fmt.Println("a manin ber:")
	fmt.Scanf("%d\n", &a)

	var b int
	fmt.Println("b manin ber:")
	fmt.Scanf("%d\n", &b)

	var c int
	fmt.Println("c manin ber:")
	fmt.Scanf("%d\n", &c)

	V := a * b * c

	fmt.Println("Natijesi:", V)
}
