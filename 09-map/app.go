package main

import "fmt"

func main() {
	prices, priceKeys := fruitPrices()
	fmt.Println(prices)
	fmt.Println(priceKeys)
	fmt.Println("apple:", prices["apple"])
	fmt.Println()

	delete(prices, "watermelon")
	fmt.Println(prices)
	fmt.Println(priceKeys)

	value, isPresent := prices["watermelon"]
	fmt.Println("default value:", value)
	fmt.Println(`is the key "watermelon" present:`, isPresent)
	fmt.Println()

	value2, isPresent2 := prices["mango"]
	fmt.Println("default value:", value2)
	fmt.Println(`is the key "mango" present:`, isPresent2)
}

func fruitPrices() (map[string]int, []string) {
	prices := map[string]int{} // ~ make(map[string]int)
	prices["apple"] = 10
	prices["banana"] = 20
	prices["watermelon"] = 30

	priceKeys := []string{}
	for key := range prices {
		priceKeys = append(priceKeys, key)
	}

	return prices, priceKeys
}
