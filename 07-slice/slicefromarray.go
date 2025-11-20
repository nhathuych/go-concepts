package main

import "fmt"

func sliceFromArray() {
	primes := [6]int{1, 2, 3, 4, 5, 6}
	fmt.Println("primes:", primes)

	var s1 []int = primes[1:4]
	fmt.Println("s1:", s1)

	var s2 []int = primes[:3]
	fmt.Println("s2:", s2)

	var s3 []int = primes[3:]
	fmt.Println("s3:", s3)

	s1[0] = 200
	fmt.Println("s1:", s1)
	fmt.Println("s2:", s2)
	fmt.Println("primes:", primes)
}
