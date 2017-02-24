package main 

import (
    "fmt"
    "math"
)

type Square struct {
    side float64
}

type Circle struct {
    radius float64 
}

type Shaper interface {
    Area() float64
}

func main() {
    var areaIntf Shaper 
    sq1 := new(Square)
    sq1.side = 5

    areaIntf = sq1 
    if t, ok := areaIntf.(*Square); ok {
        fmt.Printf("The type of areaIntf is: %T\n", t)
        fmt.Println(t)
    }
    if u, ok := areaIntf.(*Circle); ok {
        fmt.Printf("The type of areaIntf is: %T\n", u)
        fmt.Println(u)
    } else {
        fmt.Println("areaIntf does not contain a variable of type Circle")
    }
}

func (sq *Square) Area() float64 {
    return sq.side * sq.side
}

func (ci *Circle) Area() float64 {
    return ci.radius * ci.radius * math.Pi
}
