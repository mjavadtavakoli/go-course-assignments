package main

import (
	"fmt"
	"math"
)

// 1️⃣ تعریف interface
type Shape interface {
	Area() float64
}

// 2️⃣ تعریف Circle
type Circle struct {
	Radius float64
}

func (c Circle) Area() float64 {
	return math.Pi * c.Radius * c.Radius
}

// 3️⃣ تعریف Rectangle
type Rectangle struct {
	Width  float64
	Height float64
}

func (r Rectangle) Area() float64 {
	return r.Width * r.Height
}

// 4️⃣ تابعی که با interface کار می‌کند
func PrintArea(s Shape) {
	fmt.Println("Area:", s.Area())
}

func main() {
	c := Circle{Radius: 5}
	r := Rectangle{Width: 4, Height: 6}

	PrintArea(c)
	PrintArea(r)
}
