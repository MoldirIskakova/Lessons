package main

import (
	"fmt"
	"math"
)

func main() {
	m := 5
	V := math.Pow(float64(m), 3)
	S := 6 * (m * m)
	fmt.Println("Volume:", V)
	fmt.Println("Area:", S)
}
