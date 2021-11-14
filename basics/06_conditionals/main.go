package main

import "fmt"

func main() {

	x := 3
	y := 6

	if x <= y {
		fmt.Printf("%d is equal to %d or smaller than %d\n", x, y, y)
	} else {
		fmt.Printf("%d is bigger than %d\n", y, x)
	}

	color := "red"

	if color == "black" {
		fmt.Println("Color is black")
	} else if color == "white" {
		fmt.Println("Color is white")
	} else {
		fmt.Println("Color is not black or white")
	}

	switch color {
	case "black":
		fmt.Println("Color is black")
	case "white":
		fmt.Println("Color is white")
	default:
		fmt.Println("Color is not black or white")
	}

}
