package main

import (
	"fmt"
	"math/rand"
)

func main() {
	x := 10
	y := x
	x = 20
	y = 30
	fmt.Println("Value of x:", x)
	fmt.Println("Value of y:", y)
	fmt.Println("---------------------------------")
	fmt.Println()

	k := 60
	var j *int // declare a pointer to an int
	j = &k
	fmt.Println("Value of j (address of k):", j)
	fmt.Println("Value pointed to by j:", *j)
	fmt.Println("Address of j:", &j)
	fmt.Println("---------------------------------")
	fmt.Println()

	i := 42
	p := &i
	fmt.Println("Value of p:", *p)
	fmt.Println("Address of i:", p)

	*p = 23
	fmt.Println("New value of i:", i)
	fmt.Println("---------------------------------")
	fmt.Println()

	arr := []int{}
	fmt.Println("Array before:", arr)
	addRandomItem(&arr)
	addRandomItem(&arr)
	addRandomItem(&arr)
	fmt.Println("Array after:", arr)
}

func addRandomItem(arr *[]int) {
	*arr = append(*arr, rand.Intn(10000))
}
