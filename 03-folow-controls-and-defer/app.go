package main

import "fmt"

func main() {
	printOddNumbers()
	fmt.Println()

	printEvenNumbers()
	fmt.Println()

	daysUntilSaturday()
	fmt.Println()

	printGreeting()
	fmt.Println()

	fmt.Println("------------ Defer ------------")
	fmt.Println()

	testDeferOrder()
	fmt.Println()

	readFile()
	fmt.Println()
}
