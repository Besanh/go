package main

import (
	"fmt"
	"math"
	"strings"

	"golang.org/x/tour/wc"
)

type Vertex struct {
	Lat, Long float64
}

var m map[string]Vertex
var m1 = map[string]Vertex{
	"A": Vertex{
		40, 45,
	},
	"B": Vertex{
		10, 100,
	},
}

func Pic(dx, dy int) [][]uint8 {
	ar := make([][]uint8, dy)
	for x := 0; x < dx; x++ {
		a1 := make([]uint8, dx)
		for y := 0; y < dy; y++ {
			a1[y] = uint8(x * y)
		}
		ar[x] = a1
	}

	return ar
}

func WordCount(s string) map[string]int {
	x := strings.Fields(s)
	wordCount := make(map[string]int)
	for i := range x {
		wordCount[x[i]]++
	}
	return wordCount
}

// Anonymous function
func compute(fn func(float64, float64) float64) float64 {
	return fn(3, 4)
}

// Closure function
func adder() func(int) int {
	sum := 0
	return func(i int) int {
		sum += i
		return sum
	}
}

func main() {
	// pic.Show(Pic)
	m = make(map[string]Vertex)
	m["Ring Rang"] = Vertex{40.68433, -74.39967}

	// fmt.Println(m["Ring Rang"])
	// fmt.Println(m1)

	m2 := make(map[string]int)
	m2["A"] = 11
	fmt.Println(m2["A"])
	m2["A"] = 15
	fmt.Println(m2["A"])
	delete(m2, "A")
	fmt.Println(m2["A"])
	v, ok := m2["A"]
	fmt.Println("V ", v, " OK ", ok)

	wc.Test(WordCount)

	hypot := func(x, y float64) float64 {
		return math.Sqrt(x*x + y*y)
	}
	fmt.Println(hypot(2, 4))
	fmt.Println(compute(hypot))
	fmt.Println(compute(math.Pow))

	pos, neg := adder(), adder()
	for i := 0; i < 10; i++ {
		fmt.Println(pos(i), neg(-2*i))
	}
}
