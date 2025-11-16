package main

import "fmt"

func printOddNumbers() {
	n := 100
	for i := 0; i < n; i++ {
		if i%2 == 0 {
			continue
		}
		if i >= 10 {
			break // Exit the loop
		}
		fmt.Printf("%v ", i)
	}
}

func printEvenNumbers() {
	nums := []int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10}
	for i, v := range nums {
		if v%2 != 0 {
			continue
		}
		fmt.Printf("Index: %v, Value: %v\n", i, v)
	}
}
