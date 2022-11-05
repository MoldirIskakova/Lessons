package main

import "fmt"

func main() {
	x := 10
	p := &x

	fmt.Println("Значение х:", x)
	fmt.Println("Адрес х:", p)
	fmt.Println("Значение p:", *p)

	*p = 15
	fmt.Println("Значение х после изменения *p:", x)
}
