package main

import "fmt"

func main() {
	p1 := Person{
		FirstName: "Huy", // LastName will be ""
		Age:       23,
		Address: Address{
			Street: "Xóm Cải",
			City:   "Chợ Lớn",
		},
	}
	fmt.Println(p1.Introduce())
	fmt.Printf("Address: %s", p1.Address.FullAddress())
	fmt.Println()

	p2 := Person{}
	p2.FirstName = "Elon"
	p2.LastName = "Musk"
	p2.Age = 26
	p2.Address = Address{
		Street: "Rocket Road",
		City:   "Mars",
	}
	fmt.Println(p2.Introduce())
	fmt.Printf("Address: %s", p2.Address.FullAddress())
	fmt.Println()
}
