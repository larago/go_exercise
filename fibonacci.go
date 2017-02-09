package main

import "fmt"

func fibonacii(n int) (res int) {
	if n <= 1 {
		res = 1
	} else {
		res = fibonacii(n-1) + fibonacii(n-2)
	}
	return
}

func main() {
	result := 0
	for i := 0; i <= 10; i++ {
		result = fibonacii(i)
		fmt.Printf("fibonacci(%d) is %d\n", i, result)
	}
}
