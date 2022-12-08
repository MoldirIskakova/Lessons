package main

import "fmt"

func main() {
	var todoList = [3]string{
		"купить хлеб",
		"купить молоко",
		"купить пиво",
	}
	fmt.Printf("Кол-во элементов в списке: %d\n", len(todoList))
	for _, item := range todoList {
		fmt.Printf("%s\n", item)
	}

}
