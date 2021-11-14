package main

import "fmt"

func main() {

	var numbers [3]int
	numbers[0] = 10
	numbers[1] = 20
	numbers[2] = 30
	fmt.Println(numbers)
	fmt.Println(numbers[1:2])

	//decleare and assign
	names := [1]string{"Akif"}
	fmt.Println(names)

	// slices
	cities := []string{"Istanbul", "Ankara"}
	fmt.Println(cities)

}
