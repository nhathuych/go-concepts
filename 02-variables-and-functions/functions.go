package main

func addV1(a int, b int) int {
	return a + b
}

// a and b have the same type int, so we can shorten the declaration
func addV2(a, b int) int {
	return a + b
}

// swap returns two strings in swapped order
func swap(x, y string) (string, string) {
	return y, x
}
