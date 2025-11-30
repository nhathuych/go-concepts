package main

func main() {
	c := Circle{Radius: 5}
	println("Area of circle:", c.Area())

	var s Shape = Circle{Radius: 5}
	println("Area via interface:", s.Area())
}
