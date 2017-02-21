package main

import "fmt"

type TwoInts struct {
	a int
	b int
}

func main() {
	two1 := new(TwoInts)
	two1.a = 12
	two1.b = 10

	fmt.Println(two1.AddThem())
	fmt.Println(two1.AddToParam(20))

	two2 := TwoInts{3, 4}
	fmt.Println(two2.AddThem())
}

func (this *TwoInts) AddThem() int {
	return this.a + this.b
}

func (this *TwoInts) AddToParam(param int) int {
	return param + this.a + this.b
}
