package main

import (
	"fmt"
)

type struct1 struct {
	i1  int
	f1  float32
	str string
}

func main() {
	ms := new(struct1)
	ms.i1 = 10
	ms.f1 = 10.5
	ms.str = "Chris"

	fmt.Println(ms.i1)
	fmt.Println(ms.f1)
	fmt.Println(ms.str)
	fmt.Println(ms)
}
