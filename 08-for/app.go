package main

import "fmt"

func main() {
	var pow1 = []int{1, 2, 4, 8, 16, 32, 64}
	for idx, val := range pow1 {
		fmt.Printf("2**%d = %d\n", idx, val)
	}
	fmt.Println()

	pow2 := make([]int, 9)
	fmt.Println(pow2)
	for i := range pow2 {
		pow2[i] = 1 << uint(i) // 2**i
	}
	for _, value := range pow2 {
		fmt.Println(value)
	}
}
