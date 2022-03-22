package main

import (
	"fmt"
	"math"
)

// Struct
type Vertex struct {
	X, Y float64
}

// None-struct
type MyFloat float64

type Abser interface {
	Abs() float64
}

// Co the truy cap func Abs tu "v"
func (v *Vertex) Abs() float64 {
	return math.Sqrt(v.X*v.X + v.Y*v.Y)
}

func Abs1(v Vertex) float64 {
	return math.Sqrt(v.X*v.X + v.Y*v.Y)
}

func main() {
	v := Vertex{3, 4}
	fmt.Println("Method ", v.Abs())
	fmt.Println("Method with argument is function ", Abs1(v))

	f := MyFloat(-math.Sqrt2)
	fmt.Println(f.Abs2())

	v.Scale(10)
	fmt.Println(v.Abs())

	// Methods and pointer indirection
	a := Vertex{3, 4}
	v.Scale(10) // = &v.Scale(10)
	ScaleFunc(&v, 10)
	p := &Vertex{4, 3} // dia chi
	p.Scale(3)
	ScaleFunc(p, 8) // value
	fmt.Println(a, p)

	p1 := &Vertex{4, 3}
	fmt.Println(p1.Abs())
	// Dung &p la sai vi "v" argument khong phai pointer
	// Khong dung p1 ma la *p1 vi p1 dang la con tro
	fmt.Println(ScaleFuncWithoutPointer(*p1))

	// Interface
	var a1 Abser
	// f1 := MyFloat(-math.Sqrt2)
	v1 := Vertex{3, 4}
	// a1 = f1  // a1 MyFloat implement Abser
	a1 = &v1 // a1 Vertex implement Abser

	a1 = &v1
	fmt.Println("Interface ", a1.Abs())

}

func (f MyFloat) Abs2() float64 {
	if f < 0 {
		return float64(-f)
	}
	return float64(f)
}

func (v *Vertex) Scale(f float64) {
	v.X = v.X * f
	v.Y = v.Y * f
}

func ScaleFunc(v *Vertex, f float64) {
	v.X = v.X * f
	v.Y = v.Y * f
}

func ScaleFuncWithoutPointer(v Vertex) float64 {
	return math.Sqrt(v.X*v.X + v.Y*v.Y)
}
