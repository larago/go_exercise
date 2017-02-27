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

func (sq *Square) Area() float64 {
    return sq.side * sq.side 
}

func (ci *Circle) Area() float64 {
    return ci.radius * ci.radius * math.Pi
}

func main() {
    var areaIntf Shaper 
    sq1 := new(Square)
    sq1.side = 5
    areaIntf = sq1

<<<<<<< HEAD
    if t, ok := areaIntf.(*Square); ok {
        fmt.Printf("The type of areaIntf is: %T\n", t)
    }
    if u, ok := areaIntf.(*Circle); ok {
        fmt.Printf("The type of areaIntf is: %T\n", u)
    }
=======
    areaIntf = sq1 
    // if t, ok := areaIntf.(*Square); ok {
    //     fmt.Printf("The type of areaIntf is: %T\n", t)
    //     fmt.Println(t)
    // }
    // if u, ok := areaIntf.(*Circle); ok {
    //     fmt.Printf("The type of areaIntf is: %T\n", u)
    //     fmt.Println(u)
    // } else {
    //     fmt.Println("areaIntf does not contain a variable of type Circle")
    // }
    switch t := areaIntf.(type) {
        case *Square:
        fmt.Printf("Type Square %T with value %v\n", t, t)
        case *Circle:
        fmt.Printf("Type Circle %T with value %v\n", t, t)
        case nil:
        fmt.Println("nil value: nothing to check?\n")
        default:
        fmt.Printf("Unexpected type %T\n", t)
    }
    classifier(areaIntf)
}

func (sq *Square) Area() float64 {
    return sq.side * sq.side
}

func (ci *Circle) Area() float64 {
    return ci.radius * ci.radius * math.Pi
}

func classifier(items...interface{}) {
    for i, x := range items {
        switch x.(type) {
            case bool:
            fmt.Printf("Param #%d is a bool\n", i)
            case float64:
            fmt.Printf("Param #%d is a float64", i)
            case int, int64:
            fmt.Printf("Param #%d is a int\n", i)
            case nil:
            fmt.Printf("Param #%d is a nil\n", i)
            case string:
            fmt.Printf("Param #%d is a string\n", i)
            case *Square:
            fmt.Println("Parma #%s is Square")
            default:
            fmt.Printf("Param #%d is unknown\n", i)
        }
    }
>>>>>>> 7244019d1e8874ead53ddbf97ad05ead4e02c3f5
}