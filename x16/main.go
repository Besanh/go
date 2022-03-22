package main

import (
	"fmt"
	"math"
)

type Vertext struct {
	X, Y float64
}

// None struct
type MyFloat float64

// Method
func (v Vertext) Abs() float64 {
	return math.Sqrt(v.X*v.X + v.Y*v.Y)
}

// Method
func Abs2(v Vertext) float64 {
	return math.Sqrt(v.X*v.X + v.Y*v.Y)
}

func (f MyFloat) Abs3() float64 {
	if f < 0 {
		return float64(-f)
	}
	return float64(f)
}

// Pointer receiver
func (v *Vertext) Scale(f float64) {
	v.X = v.X * f
	v.Y = v.Y * f
}

func Abs5(v Vertext) float64 {
	return math.Sqrt(v.X*v.X + v.Y*v.Y)
}
func Scale1(v *Vertext, f float64) {
	v.X = v.X * f
	v.Y = v.Y * f
}

func main() {
	v := Vertext{1, 2}
	fmt.Println(v.Abs())
	fmt.Println(Abs2(v))

	// None struct
	f := MyFloat(math.Sqrt2)
	fmt.Println(f)

	v1 := Vertext{3, 4}
	v1.Scale(10) // khong the su dung Ab2() vi no chi nhan arg khong phai receiver arg
	fmt.Println(v1.Abs())

	v2 := Vertext{5, 6}
	Scale1(&v2, 10)
	fmt.Println(Abs5(v))
}
