package main

import "fmt"

func main() {
	counter := Counter{Value: 10}

	counter.IncrementByValue()
	fmt.Println("Outside IncrementByValue method:", counter.Value)

	counter.IncrementByPointer()
	fmt.Println("Outside IncrementByPointer method:", counter.Value)
}
