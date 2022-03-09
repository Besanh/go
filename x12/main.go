package main

import (
	"fmt"
	"math"
	"math/cmplx"
	"math/rand"
	"runtime"
	"time"
)

var c, java, php bool

var (
	ToBe   bool       = false
	MaxInt uint64     = 1<<64 - 1
	z      complex128 = cmplx.Sqrt(-5 + 25i)
)

const Pi = 3.14

// dich bit
const (
	// Create a huge number by shifting a 1 bit left 100 places.
	// In other words, the binary number that is 1 followed by 100 zeroes.
	Big = 1 << 100
	// Shift it right again 99 places, so we end up with 1<<1, or 2.
	Small = Big >> 99
)

func init() {
	rand.Seed(time.Now().UnixNano())
}

func main() {
	fmt.Println("Welcome to playroud!")
	fmt.Println("The time is ", time.Now())
	fmt.Println("Random number is ", rand.Intn(100))
	fmt.Println("Can bac 2 ", math.Sqrt(25))
	fmt.Println("Nextafter ", math.Nextafter(10.5, 11))
	fmt.Println(math.Pi) // pi - chu "p" viet thuong la chua dc export => loi
	fmt.Println(add(3, 2))
	fmt.Println(add(1, 10))
	fmt.Println(swap("c", "z"))
	fmt.Println(split(10))

	var i, j, k = true, 0, "ok"
	fmt.Println(i, j, k, c, java, php)

	fmt.Printf("Type %T value %v\n", ToBe, ToBe)
	fmt.Printf("Type %T value %v\n", MaxInt, MaxInt)
	fmt.Printf("Type %T value %v\n", z, z)

	var ii int
	var f float64
	var b bool
	var s string
	fmt.Printf("%v %v %v %q\n", ii, f, b, s)

	// Ep kieu data
	var x, y int = 3, 4
	var ff float64 = math.Sqrt(float64(x*x - y*y))
	var z uint = uint(ff)
	fmt.Printf("%v, %v, %v, %v\n", x, y, ff, z)
	fmt.Println("Pi ", Pi, " day\n")

	// Dich bit
	fmt.Println(needInt(Small), "\n")
	fmt.Println(needFloat(Small), "\n")
	fmt.Println(needFloat(Big), "\n")

	// For loop
	total := 0
	for i := 0; i < 10; i++ {
		total += i
	}
	fmt.Println("For loop ", total, "\n")
	sum := 1
	for sum < 100 {
		sum += sum
	}
	fmt.Println("Sort for loop %v\n", sum)

	fmt.Println("Use if ", sqrt(2), sqrt(-4), "\n")

	fmt.Println("Luy thua ", pow(2, 3, 10))
	fmt.Println("Lim ", pow(12, 3, 10))

	// Switch case
	switch os := runtime.GOOS; os {
	case "darwin":
		fmt.Println("OS X\n")
	case "linux":
		fmt.Println("OS Linux\n")
	default:
		fmt.Printf("OS %s\n", os)
	}

	// Switch case evalution order
	fmt.Println("When's Saturday?\n")
	today := time.Now().Weekday()
	fmt.Println(today, "\n")
	switch time.Saturday {
	case today + 0:
		fmt.Print("Today\n")
	case today + 1:
		fmt.Print("Tomorrow\n")
	case today + 2:
		fmt.Print("In 2 days\n")
	default:
		fmt.Print("Too far away\n")
	}

	// Switch case without condition
	t := time.Now()
	switch {
	case t.Hour() < 12:
		fmt.Println("Good morning\n")
	case t.Hour() < 17:
		fmt.Println("Good afternoon\n")
	default:
		fmt.Println("Good evening\n")
	}

	// Defer
	defer fmt.Println("World\n")
	fmt.Println("hello")
	// Stacking defer
	fmt.Println("Stacking defer ")
	for i := 0; i < 10; i++ {
		defer fmt.Println(i)
	}
	fmt.Println("done")
}

func add(x int, y int) int {
	return x + y
}

// Rut gon type
func add1(x, y int) int {
	return x + y
}

func swap(a, b string) (string, string) {
	return b, a
}

// Naked return
func split(sum int) (x, y int) {
	x = sum * 4 / 9
	y = sum - x
	return
}

func needInt(x int) int { return x*10 + 1 }
func needFloat(x float64) float64 {
	return x * 0.1
}

func sqrt(x float64) string {
	if x < 0 {
		return sqrt(-x) + "i"
	}
	return fmt.Sprint(math.Sqrt(x))
}

// Tinh luy thua
func pow(x, n, lim float64) float64 {
	if v := math.Pow(x, n); v < lim {
		return v
	}
	return lim
}
