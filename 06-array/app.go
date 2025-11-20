package main

import "fmt"

func main() {
	var a [2]string
	fmt.Println(a[0])
	fmt.Println(a[1])
	fmt.Println(len(a))

	a[0] = "Hello"
	a[1] = "World"
	fmt.Println(a)
}
