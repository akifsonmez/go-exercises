package main

import "fmt" // short form of format

func main() {

	fmt.Printf("Hello ")
	fmt.Printf("World\n")
	fmt.Println("Hello", 25, "Akif")

	myName := returnMyName()
	sayMyName(myName)
}

func sayMyName(name string) {
	fmt.Println(name)
}

func returnMyName() string {
	return "Akif"
}
