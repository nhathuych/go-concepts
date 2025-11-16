package main

import "fmt"

func declareVariables() {
	// Declare a variable with var keyword and explicit type
	var a, b int = 10, 20
	fmt.Println("a:", a, "b:", b)

	var c int8 = 127 // int8 has a smaller range than int
	fmt.Println("c:", c+1)
}

func defaultValues() {
	var i int
	var f float64
	var b bool
	var s string

	fmt.Println("Default values:")
	fmt.Printf("%v %v %v %q\n", i, f, b, s)
}

func castVariables() {
	var i int = 69
	var f float64 = float64(i)
	var u float64 = 9223372036854775807.0
	var j int32 = int32(u)

	fmt.Println("Casting variables:")
	fmt.Printf("i: %v, f: %v, u: %v, j: %v\n", i, f, u, j)
}

func constants() {
	const pi = 3.14159 // unexported constant
	const Pi = 3.14    // exported constant
	// Pi += 0.00159 // This will cause a compile-time error if uncommented
	const (
		StatusOK       = 200
		StatusNotFound = 404
	)

	fmt.Println("Constants:")
	fmt.Printf("unexported pi: %v\n", pi)
	fmt.Printf("exported Pi: %v\n", Pi)
	fmt.Printf("StatusOK: %v, StatusNotFound: %v\n", StatusOK, StatusNotFound)
}
