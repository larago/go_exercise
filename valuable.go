package main 

import (
    "fmt"
    "reflect"
)

type stockPosition struct {
    ticker      string 
    shareProce  float32
    count       float32
}

func (s stockPosition) getValue() float32 {
    return s.shareProce * s.count
}

type car struct {
    make    string 
    model   string
    price   float32 
}

func (c car) getValue() float32 {
    return c.price 
}

type valuable interface {
    getValue()  float32
}

func showValue(asset valuable) {
    fmt.Printf("Value of the asset is: %f\n", asset.getValue())
}

func main() {
    var o valuable = stockPosition{"GOOG", 577.20, 4}
    oType := reflect.TypeOf(o)
    fmt.Println(oType)
    showValue(o)
    o = car{"BMW", "M3", 66500}
    showValue(o)
    oType = reflect.TypeOf(o)
    fmt.Println(oType)
}


