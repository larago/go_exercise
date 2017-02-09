package main

import (
	"fmt"
	"time"
)

func main() {
	result := 0
	start := time.Now()
	for i := 0; i <= 50; i++ {
		result = fibonacii(i)
		fmt.Printf("fibonacci(%d) is: %d\n", i, result)
	}
	end := time.Now()
	delta := end.Sub(start)
	fmt.Printf("longCalculation took this amount of time: %s\n", delta)
}

func fibonacii(n int) (res int) {
	if n <= 1 {
		res = 1
	} else {
		res = fibonacii(n-1) + fibonacii(n-2)
	}
	return
}
