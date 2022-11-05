package main

import (
	"errors"
	"fmt"
)

const Pi = 3.1415

func main() {
	printCircleArea(-2)
}

func printCircleArea(radius int) {
	circleArea, err := calculateCircleArea(radius)
	if err != nil {
		fmt.Println(err.Error())
		return
	}
	fmt.Printf("Радиус круга: %d см \n", radius)
	fmt.Println("Формула для расчета площади круга: A=πr2\n ")

	fmt.Printf("площадь круга: %f32 см. кв.\n", circleArea)
}

func calculateCircleArea(radius int) (float32, error) {
	if radius <= 0 {
		return float32(0), errors.New("Радиус круга не может быть отрицательным!")
	}

	return float32(radius) * float32(radius) * Pi, nil
}
