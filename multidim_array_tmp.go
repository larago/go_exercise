package main

import "fmt"

func main() {
	var screen [][]int = make([][]int, 10)
	for row := range screen {
		for column := range screen[row] {
			screen[row][column] = 1
		}
	}

	fmt.Println(screen)
}
