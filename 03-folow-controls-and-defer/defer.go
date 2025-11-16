package main

import (
	"fmt"
	"os"
)

func testDeferOrder() {
	fmt.Println("A")
	defer fmt.Println("B")
	defer fmt.Println("C")
	fmt.Println("D")
}

func readFile() {
	file, error := os.Open("test.txt")
	if error != nil {
		fmt.Println("Error opening file:", error)
		return
	}

	// always executed, even when returning early
	defer file.Close()

	fmt.Println("File opened successfully")
}
