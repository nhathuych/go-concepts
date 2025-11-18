package main

import "fmt"

type Address struct {
	Street string
	City   string
}

func (a Address) FullAddress() string {
	return fmt.Sprintf("%s, %s", a.Street, a.City)
}
