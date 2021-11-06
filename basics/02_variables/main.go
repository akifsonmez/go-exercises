package main

import "fmt"

func main() {

	var name string = "Akif"
	var surname = "Sonmez" // type inference
	// var age int32 = 25
	age := 25 // shorthand // cannot be used for global variables
	isGoFun := true

	fmt.Println(name, surname, age)
	fmt.Printf("%T\n", age)
	fmt.Printf("%T\n", name)
	fmt.Printf("%T\n", isGoFun)

}
