package main 

import (
	"fmt"
)

func main() {
	aMap := make(map[string]string)
	aMap["kk"] = "qq"
	fmt.Println(aMap["aa"])
	fmt.Println(len(aMap["aa"]))
}