package main

import (
	"fmt"
	"math"
)

type Rectangular struct {
	width,height float64
}

type Circle struct {
	diameter float64
}
func main() {
	r := Rectangular{2,3}
	area1 := r.Area()
	fmt.Println(area1)
	r.setValue1()
	fmt.Println(r.Area())
	//r.setValue2()
	fmt.Println(r.Area())

	c := Circle{1.381977}
	area2 := c.Area()
	fmt.Println(area2)
}

func (r Rectangular) Area()float64 {
	return  r.width * r.height
}

func (c Circle) Area()float64  {
	return math.Pi * c.diameter * c.diameter
}

func (r Rectangular) setValue1()  {
	r.height = 5
}

