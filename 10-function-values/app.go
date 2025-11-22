package main

import "fmt"

func main() {
	fn1 := func(a, b int) int {
		return a - b
	}
	fmt.Println("fn1:", fn1(3, 2))
	fmt.Println()

	fn2 := add
	fmt.Println("fn2:", fn2(2, 4))
	fmt.Println()

	fmt.Println("apply fn1:", apply(fn1))
	fmt.Println("apply fn2:", apply(fn2))
}

func add(a, b int) int {
	return a + b
}

func apply(f func(a, b int) int) int {
	return f(6, 9)
}
