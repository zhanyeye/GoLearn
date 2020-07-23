package structsmethodsinterfaces

import "math"

// This is quite different to interfaces in most other programming languages.
// Normally you have to write code to say My type Foo implements interface Bar.
// But in our case
// - Rectangle has a method called Area that returns a float64 so it satisfies the Shape interface
// - Circle has a method called Area that returns a float64 so it satisfies the Shape interface
// - string does not have such a method, so it doesn't satisfy the interface
// - etc.

type Shape interface {
	Area() float64
}


type Rectangle struct {
	Width float64
	Height float64
}

func Perimeter(rectangle Rectangle) float64 {
	return 2 * (rectangle.Height + rectangle.Width)
}

func Area(rectangle Rectangle) float64 {
	return rectangle.Height * rectangle.Width
}

// next requirement is to write an Area function for circles.
// GO not support overloading of methods and operators.
// - we can have functions with the same name declared in different packages.
// - We can define methods on our newly defined types instead.

type Circle struct {
	Radius float64
}

// The only difference between methods and functions is the syntax of the method receiver
// func (receiverName ReceiverType) MethodName(args)
// When your method is called on a variable of that type, you get your reference to its data via the receiverName variable
// In many other programming languages this is done implicitly and you access the receiver via this.
// It is a convention in Go to have the receiver variable be the first letter of the type.
func (c Circle) Area() float64 {
	return math.Pi * c.Radius * c.Radius
}

func (r Rectangle) Area() float64 {
	return r.Width * r.Height
}

type Triangle struct {
	Base float64
	Height float64
}

func (t Triangle) Area() float64 {
	return 0.5 * t.Height * t.Base
}