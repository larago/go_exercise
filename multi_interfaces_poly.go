package main 

import "fmt"

type Shaper interface {
    Area() float64
}

type TopologicalGenus interface {
    Rank() int
}

type Square struct {
    side float64 
}

func (sq *Square) Area() float64 {
    return sq.side * sq.side 
}

func (sq *Square) Rank() int {
    return 1
}

type Rectangle struct {
    length, width float64 
}

func (r Rectangle) Area() float64 {
    return r.length * r.length
}

func (r Rectangle) Rank() int {
    return 2
}

func main() {
    r := Rectangle{5, 3}
    q := &Square{5}
    shapes := []Shaper{r, q}
    fmt.Println("Looping through shapes for area...")
    for n, _ := range shapes {
        fmt.Println("Shape details:", shapes[n])
        fmt.Println("Area of this shape is:", shapes[n].Area())
    }
    topgen := []TopologicalGenus{r, q}
    fmt.Println("Looping through topgen for rank...")
    for n, _ := range shapes {
        fmt.Println("Shape details:", topgen[n])
        fmt.Println("Topological Genus of this shape is:", topgen[n].Rank())
    }
}