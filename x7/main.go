package main

import "fmt"

type Number interface {
	int64|float64
}

func main() {
	ints := map[string]int64{
		"first":  1,
		"second": 2,
	}

	floats := map[string]float64{
		"first":  1.5,
		"second": 2.1,
	}

	fmt.Printf("Non-generic: %v and %v\n",
		NonGenericInt(ints),
		NonGenericFloat(floats))
	fmt.Printf("Generic %v and %v\n",
		SumIntOrFloat(ints),
		SumIntOrFloat(floats))
	fmt.Printf("Generic without parameter: %v and %v\n",
		SumNumber(ints),
		SumNumber(floats))
}

// Non generic int
// map[string]int64 tra ve cap key-value dinh dang kieu string-int64
func NonGenericInt(m map[string]int64) int64 {
	var s int64
	for _, v := range m {
		s += v
	}
	return s
}

// None generic float
func NonGenericFloat(m map[string]float64) float64 {
	var s float64
	for _, v := range m {
		s += v
	}
	return s
}

// Generic with argument
// s type la V(int64|float64)
func SumIntOrFloat[K comparable, V int64 | float64](m map[K]V) V {
	var s V
	for _, v := range m {
		s += v
	}
	return s
}

// Generic without argument
func SumNumber[K comparable, V Number](m map[K]V) V {
	var s V
	for _, v := range m {
		s += v
	}
	return s
}