package main

import "fmt"

type Counter struct {
	Value int
}

// IncrementByValue increases the counter by 1,
// but since the receiver is a value (a copy of Counter),
// the change does not affect the original variable.
func (c Counter) IncrementByValue() {
	c.Value++
	fmt.Println("Inside IncrementByValue method:", c.Value)
}

// IncrementByPointer increases the counter by 1,
// and since the receiver is a pointer (*Counter),
// the change modifies the original variable.
func (c *Counter) IncrementByPointer() {
	c.Value++
}
