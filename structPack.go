package main

import (
	"./struct_pack/structPack"
	"fmt"
)

func main() {
	struct1 := new(structPack.ExpStruct)
	struct1.Mi1 = 10
	struct1.Mf1 = 16.

	fmt.Println(struct1.Mi1)
	fmt.Println(struct1.Mf1)
}
