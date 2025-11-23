package main

import "fmt"

func main() {
	positive := adder()
	negative := adder()

	for i := 0; i < 10; i++ {
		fmt.Println(
			positive(i),
			negative(-2*i),
		)
	}
}

// adder returns a closure: a function that captures
// and remembers the variable "sum" even after adder() has returned.
func adder() func(int) int {
	sum := 0 // this variable is captured by the returned function
	return func(x int) int {
		// The closure updates and returns the captured "sum"
		sum += x
		return sum
	}
}
