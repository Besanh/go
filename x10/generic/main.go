package main

import (
	"fmt"
)

type Number interface {
	int64 | float64
}

func main() {
	var inputInts = map[string]int64{
		"mot": 1,
		"hai": 3,
	}

	var inputFloats = map[string]float64{
		"ba":  1.5,
		"bon": 2.3,
	}

	fmt.Printf("Non generic - Sum int va float lan luot la : %v, %v\n", sumInts(inputInts), sumFloats(inputFloats))
	fmt.Printf("Generic with param - Sum int and float: %v, %v\n", SumIntOrFloat[string, int64](inputInts), SumIntOrFloat[string, float64](inputFloats))
	fmt.Printf("Generic without param - Sum int and float: %v, %v\n", SumIntOrFloat(inputInts), SumIntOrFloat(inputFloats))
	fmt.Printf("Generic with Interface: %v, %v\n", SumNumber(inputInts), SumNumber(inputFloats))
}

func sumInts(m map[string]int64) int64 {
	var s int64
	for _, i := range m {
		s += i
	}
	return s
}

func sumFloats(m map[string]float64) float64 {
	var s float64
	for _, i := range m {
		s += i
	}
	return s
}

func SumIntOrFloat[K comparable, V int64|float64](m map[K]V) V {
	var s V
	for _, i := range m {
		s += i
	}
	return s
}
func SumNumber[K comparable, V Number](m map[K]V) V {
	var s Number
	for _, i := range m {
		s += i
	}
	return s
}