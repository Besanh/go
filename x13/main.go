package main

import (
	"fmt"
	"math"

	"golang.org/x/tour/pic"

	"golang.org/x/tour/wc"
)

type Vertext struct {
	X, Y int
}

type Vertu struct {
	Lat, Long float64
}

var m map[string]Vertu

// Struct literal
var (
	v1 = Vertext{1, 3}
	v2 = Vertext{X: 20} // Y=0
	v3 = Vertext{}      // X=Y=0
	v4 = &Vertext{1, 3} // type *Vertext
)

func main() {
	i, j := 10, 43

	p := &i // gan p dia chi cua i

	fmt.Println(*p) // print value p
	*p = 21
	fmt.Println(i) // =21, vi p su dung con tro cua i

	p = &j
	*p = *p / 3
	fmt.Println(j)

	// Struct
	fmt.Println("Struct ", Vertext{1, 3})
	//  Struct field, access struct
	v := Vertext{2, 5}
	v.X = 4
	fmt.Printf("Access data in struct %v, %v\n", v.X, v.Y)

	// Con tro den struct
	a := Vertext{3, 4}
	b := &a // kieu khac (*b).a
	b.X = 1e9
	fmt.Println("Con tro den struct ", a)

	// Struct literal
	fmt.Println("Struct literal", v1, v2, v3, v4, v4.Y)

	// Array
	primes := [3]int{1, 3, 4}
	fmt.Println(primes)
	aa := "oka"
	fmt.Println(aa[1])

	// Slice
	primess := [6]int{2, 3, 5, 7, 11, 13}
	var s []int = primess[1:4]
	fmt.Println(s)

	// Slicde references to array
	c := [4]string{
		"John",
		"Naul",
		"Carrot",
		"Luffy",
	}
	fmt.Println(c)
	d := c[0:2]
	e := c[1:4]
	fmt.Println(d, e)
	d[0] = "Zoro"
	fmt.Println(d, e)
	fmt.Println(c)

	//  Slice literal
	q := []int{2, 3, 5, 7, 11, 13} // slice
	fmt.Println(q)
	g := []struct {
		h int
		j bool
	}{
		{2, false},
		{3, true},
		{10, true},
		{101, true},
	}
	fmt.Println(g)

	// Slice default
	aaa := []int{2, 3, 4, 5, 7, 10, 12}
	fmt.Println(aaa[1:2])
	fmt.Println(aaa[:3])
	fmt.Println(aaa[1:])

	// Len and capcity slice
	ss := []int{5, 4, 6, 7, 10, 12}
	printSlice(ss) // 7:7
	ss = ss[:0]    // lay 0 phan tu
	printSlice(ss) // 0:7
	ss = ss[:4]    // la 4 phan tu dau tien
	printSlice(ss) //	5 4 6 7
	ss = ss[2:]    // 2:0
	printSlice(ss) // lay tu phan tu thu 2, lay 0 phan tu
	var k []int
	fmt.Println(k, len(k), cap(k))
	if s == nil {
		fmt.Println("nil!")
	}

	// Create slice with make
	u1 := make([]int, 5)
	printSlice1("u1", u1) // size 5, cap 5
	u2 := make([]int, 0, 5)
	printSlice1("u2", u2) // size 0, cap 5
	u3 := u2[:2]
	printSlice1("u3", u3) // size 2, len 5
	u4 := u3[2:5]
	printSlice1("u4", u4)

	// Range continued
	pow := make([]int, 10)
	for i := range pow {
		pow[i] = 1 << uint(i) // == 2**i
	}
	for _, value := range pow {
		fmt.Printf("%d\n", value)
	}

	// Picture
	pic.Show(Pic)

	// Map
	m = make(map[string]Vertu)
	m["Lab"] = Vertu{
		40.68433, -74.39967,
	}
	fmt.Println("Map ", m["Lab"])

	// Map literal
	var mm = map[string]Vertu{
		"A": Vertu{
			40.68433, -74.39967,
		},
		"Google": Vertu{
			37.42202, -122.08408,
		},
	}
	fmt.Println("Map literal ", mm)

	// Mutating map
	b1 := make(map[string]int)
	b1["A"] = 12
	fmt.Println("Mutating ", b1["A"])
	b1["B"] = 46
	delete(b1, "B")
	fmt.Println("Mutating delete ", b1["B"])
	v1, ok1 := b1["A"]
	fmt.Println(" Mutating ", v1, "Present? ", ok1)

	wc.Test(WordCount)

	// Function value
	hypot := func(x, y float64) float64 {
		return math.Sqrt(x*x + y*y)
	}
	fmt.Println(hypot(5, 12))
	fmt.Println(compute(hypot))
	fmt.Println(compute(math.Pow))

	// Closure function
	pos, neg := adder(), adder()
	for ii := 0; ii < 10; ii++ {
		fmt.Println("Closure function ", pos(ii), neg(-2*ii))
	}

	// Fibonacci
	f := fibonacci()
	for i := 0; i < 10; i++ {
		fmt.Println("Fibonacci ", f(i))
	}

}

func printSlice(s []int) {
	fmt.Printf("len slice %d, capacity %d, %v\n", len(s), cap(s), s)
}

func printSlice1(s string, x []int) {
	fmt.Printf("%s, len=%d, cap=%d, %v\n", s, len(x), cap(x), x)
}

func Pic(dx, dy int) [][]uint8 {
	ar := make([][]uint8, dy)
	for y := 0; y < dy; y++ {
		ars := make([]uint8, dx)
		for x := 0; x < dx; x++ {
			ars[x] = uint8((x + y) / 2)
		}
		ar[y] = ars
	}
	return ar
}

func WordCount(s string) map[string]int {
	return map[string]int{"Go!": 1, "I": 1, "am": 1, "learning": 1}
}

// fn co 2 argument
// compute co 1 argument
func compute(fn func(float64, float64) float64) float64 {
	return fn(3, 4)
}

// Closure
func adder() func(int) int {
	sum := 0
	return func(i int) int {
		sum += i
		return sum
	}
}

// fibonacci is a function that returns
// a function that returns an int.
func fibonacci() func(int) int {
	a1, a2 := 1, 1
	i := 3
	var a = 0
	return func(n int) int {
		if n == 1 || n == 2 {
			return 1
		}
		for i <= n {
			a = a1 + a2
			a1 = a2
			a2 = a
			i++
		}
		return a
	}
}
