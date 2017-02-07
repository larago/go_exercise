package main

import "fmt"

func main() {
	// var num1 int = 7

	// switch {
	// case num1 < 0:
	// 	fmt.Println("Number is negative")
	// case num1 > 0 && num1 < 10:
	// 	fmt.Println("Number is betweem 0 and 10")
	// default:
	// 	fmt.Println("Number is 10 or greater")
	// }
	switch result := calculate(); {
	case result < 0:
		fmt.Println("result is less than 0")
	default:
		fmt.Println("result is equal to 0 or greater.")
	}
}

func calculate() int {
	return 7
}
