package main

import "fmt"

func sendMessage(name string) {
	fmt.Println("Hello, " + name + "!")
}

func sendMessageV2(name ...string) {
	n := "Golang"
	if len(name) > 0 {
		n = name[0]
	}
	fmt.Println("Hello, " + n + "!")
}
