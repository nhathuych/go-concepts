package main

import (
	"fmt"
	"strings"
)

type Person struct {
	FirstName string
	LastName  string
	Age       int
	Address   // Embedding struct Address
}

func (p Person) FullName() string {
	// return p.FirstName + " " + p.LastName // dont do this
	var builder strings.Builder
	builder.WriteString(p.FirstName)
	builder.WriteString(" ")
	builder.WriteString(p.LastName)
	return strings.TrimSpace(builder.String())
}

func (p Person) Introduce() string {
	return fmt.Sprintf("\nHello, my name is %s and I am %d years old.", p.FullName(), p.Age)
}
