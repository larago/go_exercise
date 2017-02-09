package main

import "fmt"

func main() {
	fmt.Printf("Multiply 2 * 5 * 6 = %d\n", Multiply(2, 5, 6))
}

func Multiply(a, b, c int) int {
	return a * b * c
}
