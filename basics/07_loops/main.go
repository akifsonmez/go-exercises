package main

import "fmt"

func main() {

	// long way
	i := 1
	for i <= 5 {
		fmt.Printf("Number %d\n", i)
		i = i + 1
	}

	// short way
	for j := 1; j <= 5; j++ {
		fmt.Printf("Number %d\n", j)
	}

	// FizzBuzz
	for i := 1; i <= 100; i++ {
		if i%15 == 0 {
			fmt.Println("FizzBuzz")
		} else if i%3 == 0 {
			fmt.Println("Fizz")
		} else if i%5 == 0 {
			fmt.Println("Buzz")
		} else {
			fmt.Println(i)
		}
	}
}
