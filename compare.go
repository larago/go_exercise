package main

import (
	"fmt"
)

func compare(a, b []byte) int {
	for i := 0; i < len(a) && i < len(b); i++ {
		switch {
		case a[i] > b[i]:
			return 1
		case a[i] < b[i]:
			return -1
		}
	}

	switch {
	case len(a) > len(b):
		return -1
	case len(a) < len(b):
		return 1
	}
	return 0
}

func main() {
	result := compare([]byte{1, 2}, []byte{3, 1})
	fmt.Println(result)
}
