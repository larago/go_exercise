package main

import "fmt"

var num int = 10
var numx2, numx3 int

func main() {
	numx2, numx3 = getx2Andx3(num)
	PrintValues()
	numx2, numx3 = getx2Andx3_2(num)
	PrintValues()
}

func PrintValues() {
	fmt.Printf("num = %d, 2x num = %d, 3x num = %d\n", num, numx2, numx3)
}

func getx2Andx3(input int) (int, int) {
	return 2 * input, 3 * input
}

func getx2Andx3_2(input int) (x2 int, x3 int) {
	x2 = 2 * input
	x3 = 3 * input
	return
}
