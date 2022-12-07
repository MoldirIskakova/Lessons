package main

import "fmt"

func main() {
	ages := make(map[string]int)
	ages["Maksim"] = 20
	ages["Oleg"] = 30
	ages["Petr"] = 40

	for key, value := range ages {
		fmt.Printf("%s-%d\n", key, value)
	}

}
