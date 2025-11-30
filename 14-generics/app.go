package main

import "fmt"

func main() {
	s := []int{1, 2, 3, 4, 5}
	x := 3
	index := Index(s, x)
	fmt.Printf("Index of %d is: %d\n", x, index)

	fruits := []string{"apple", "cherry", "banana", "orange"}
	y := "banana"
	fruitIndex := Index(fruits, y)
	fmt.Printf("Index of %s is: %d\n", y, fruitIndex)

	fruitPrices := map[string]float64{
		"apple":  1,
		"cherry": 1.5,
		"banana": 2,
		"orange": 2.5,
	}
	totalPrice := SumIntsOrFloats(fruitPrices)
	fmt.Printf("Total price of fruits: %.2f\n", totalPrice)
}

func Index[TNe comparable](s []TNe, x TNe) int {
	for i, v := range s {
		if v == x {
			return i
		}
	}
	return -1
}

func SumIntsOrFloats[KEY comparable, VALUE int64 | float64](m map[KEY]VALUE) VALUE {
	var sum VALUE
	for _, value := range m {
		sum += value
	}
	return sum
}
