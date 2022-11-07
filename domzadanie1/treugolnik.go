package main

import (
	"errors"
	"fmt"
)

func main() {
	var a int
	fmt.Println("Enter  a value : ")

	_, err := fmt.Scanf("%d", &a)

	if err != nil {
		fmt.Println(err)
	}

	var h int
	_, err = fmt.Scanf("%d", &h)

	if err != nil {
		fmt.Println(err)
	}

	result, err := calculateTriangleArea(a, h)
	if err != nil {
		fmt.Println(err.Error())
		return
	}

	fmt.Printf("Высота треугольника: %d см \n", h)
	fmt.Printf("Длина треугольника: %d см \n", a)
	fmt.Println("Формула для расчета площади треугольника: S=1/2*a*h\n ")
	fmt.Printf("площадь треугольника: %f см. кв.\n", result)
}

func calculateTriangleArea(height int, width int) (float32, error) {
	if height <= 0 {
		return float32(0), errors.New("Высота не может быть отрицательным!")
	}
	if width <= 0 {
		return float32(0), errors.New("Ширина не может быть отрицательным!")
	}
	return float32(height) * float32(width) / 2, nil
}
