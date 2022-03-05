package main

import "fmt"

// Constraint inteface
// de su dung nhieu noi
// Khai bao Number la type inteface, dung nhu constraint inteface
type Number interface {
    int64 | float64
}

// SumInts adds together the values of m.
func SumInts(m map[string]int64) int64 {
	var s int64
	for _, v := range m {
		s += v
	}

	return s
}

// SumFloats adds together the values of m.
func SumFloats(m map[string]float64) float64 {
	var s float64
	for _, v := range m {
		s += v
	}

	return s
}

func main() {
	// Initialize a map for the integer values
	ints := map[string]int64{
		"first":  34,
		"second": 12,
	}

	// Initialize a map for the float values
	floats := map[string]float64{
		"first":  12.4,
		"second": 34.5,
	}

	// Non generic
	fmt.Printf("Non-generic Sums: %v and %v\n", SumInts(ints), SumFloats(floats))
	
	// Generic with type parameter
	fmt.Printf("Generic Sums: %v and %v\n", 
	SumIntsOrFloats[string, int64](ints),
	SumIntsOrFloats[string, float64](floats))

	// Generic without type parameter
	// Compiler se infer(suy luan) type cua argument
	fmt.Printf("Generic Sums, type parameters inferred: %v and %v\n",
	SumIntsOrFloats(ints),
	SumIntsOrFloats(floats))

	// Constraint inteface
	fmt.Printf("Generic Sums with Constraint: %v and %v\n",
        SumNumbers(ints),
        SumNumbers(floats))
}

// SumIntsOrFloats sums the values of map m. It supports both int64 and float64
// as types for map values.
/**
* 2 type param K V(ben trong [])
* argument m su dung 2 type param nay, m cua type map[K]V
* K type param co kieu constraint la comparable
* comparable duoc khai bao truoc trong Go
* cho phep bat ky type nao co gia tri co the dem so sanh bang toan tu ==, !=
* co the su dung K lam key trong bien map
* V type param 1 constraint(rang buoc) cua 2 type int64 va float64
* m map[K]V, neu khong khai bao K type param comparable, thi compiler se loai bo tham chieu toi map[K]V
* @return value cua type V(int hoac float)
*/
func SumIntsOrFloats[K comparable, V int64 | float64](m map[K]V) V {
	var s V
	for _, v := range m {
		s += v
	}
	return s
}

// SumNumbers sums the values of map m. It supports both integers
// and floats as map values.
func SumNumbers[K comparable, V Number](m map[K]V) V {
	var s V
	for _, v := range m {
		s+=v
	}
	return s
}