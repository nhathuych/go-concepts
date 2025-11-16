package main

import "fmt"

func main() {
	name := "Huy"
	sendMessage(name)
	fmt.Println("")

	sendMessageV2("Huy")
	sendMessageV2("Huy", "Huy Huy")
}
