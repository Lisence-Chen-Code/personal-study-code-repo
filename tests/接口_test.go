package test

import (
	"fmt"
	"math"
	"testing"
)

func TestAssertIntf(t *testing.T) {
	var s Shape = Circle{3, "red"}
	c := s.(Circle)
	fmt.Printf("%T\n", c)
	fmt.Printf("%v\n", c)
	fmt.Println("area: ", c.Area())
	fmt.Println("perimeter: ", c.Perimeter())
	fmt.Println("color:", c.Color())
}

func TestImIntf(t *testing.T) {
	c := &Circle{
		radius: 3,
		color:  "green",
	}
	fmt.Printf("%T\n", c)
	fmt.Printf("%v\n", c)
	fmt.Println("area: ", c.Area())
	fmt.Println("perimeter: ", c.Perimeter())
	fmt.Println("color:", c.Color())
}

type Shape interface {
	Area() float32
	Color() string
}

type Object interface {
	Perimeter() float32
}

type Circle struct {
	radius float32
	color  string
}

func (c Circle) Area() float32 {
	return math.Pi * (c.radius * c.radius)
}

func (c Circle) Perimeter() float32 {
	return 2 * math.Pi * c.radius
}

func (c Circle) Color() string {
	return c.color
}
