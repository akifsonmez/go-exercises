package main

import "fmt"

func main() {
	fmt.Println(greetings("Akif"))
	fmt.Println(getSum(4, 6))
}

func greetings(name string) string {
	return "Hello " + name
}

func getSum(num1, num2 int) int {
	return num1 + num2
}
