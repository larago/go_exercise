package main 

import (
	"errors"
	"fmt"
	"math"
)

var errNotFound error = errors.New("Not found error")

type PathError struct {
	Op 		string 
	Path	string 
	Err 	error
}

func (e *PathError) String() string {
	return e.Op + e.Path + ": " + e.Err.Error()
}

func main() {
	if f, err := Sqrt(-1); err != nil {
		if e, ok := err.(*os.PathError); ok {
			fmt.Println(e)
		}
	} else {
		fmt.Printf("Sqrt(%d) is: %f\n", f)
	}
}

func Sqrt(f float64) (float64, string) {
	if f < 0 {
		return 0, 
	} else {
		return math.Sqrt(f), nil
	}
}