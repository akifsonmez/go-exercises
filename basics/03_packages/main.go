package main

import (
	"fmt"
	"go-exercises/basics/03_packages/custom_math_package"
	"math"
)

func main() {
	fmt.Println(math.Ceil(3.5))
	fmt.Println(math.Floor(3.5))
	fmt.Println(math.Sqrt(16))
	fmt.Println(custom_math_package.Square(4))
}
