package main

import "fmt"

func makeSlice() {
	a := make([]int, 3) // ~ a := []int{0, 0, 0}
	fmt.Println("a:", a)
	fmt.Println("len(a):", len(a)) // length = current size
	fmt.Println("cap(a):", cap(a)) // capacity = maximum size before reallocating
	fmt.Println()

	b := make([]int, 0, 4)
	fmt.Println("b:", b)
	fmt.Println("len(b):", len(b))
	fmt.Println("cap(b):", cap(b))
	b = append(b, 10, 20, 30, 40, 50)
	fmt.Println("b (after append):", b)
	fmt.Println("len(b):", len(b))
	fmt.Println("cap(b):", cap(b))
}
