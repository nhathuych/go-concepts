package main

import (
	"fmt"
)

func main() {
	sum1 := addV1(1, 2)
	sum2 := addV2(10, 20)

	fmt.Println("Sum1:", sum1)
	fmt.Println("Sum2:", sum2)
	fmt.Println()

	str1, str2 := swap("hello", "world")
	fmt.Println("Swapped strings:", str1, str2)
	fmt.Println()

	fmt.Println("Variable Declaration:")
	declareVariables()
	fmt.Println()

	defaultValues()
	fmt.Println()

	castVariables()
	fmt.Println()

	constants()
	fmt.Println()
}
