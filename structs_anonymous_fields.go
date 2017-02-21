package main

import "fmt"

type innerS struct {
	in1 int
	in2 int
}

type outerS struct {
	b int
	c float32
	int
	innerS
}

func main() {
	outer := new(outerS)
	outer.b = 6
	outer.c = 7.5
	outer.int = 60
	outer.in1 = 5
	outer.in2 = 10

	fmt.Println(outer.b)
	fmt.Println(outer.c)
	fmt.Println(outer.int)
	fmt.Println(outer.in1)
	fmt.Println(outer.in2)

	outer2 := outerS{6, 7.5, 60, innerS{5, 10}}
	fmt.Println(outer2)
}
