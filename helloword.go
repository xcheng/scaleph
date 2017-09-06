package main

import (
	"fmt"
	"math"
	//"math/cmplx"
	"math/rand"
	"runtime"
	"time"
)

var (
	ToBe   bool   = false
	MaxInt uint64 = 1<<64 - 1
	//z                 complex128 = cmplx.Sqrt(-5 + 12i)
	java, python, ios     = true, true, "No"
	i, j              int = 3, 4
)

func main() {

	test()
	a, b := swap("Hello", "World")
	fmt.Println(a, b)
	flowFor()
	fmt.Println(sqrt(2), sqrt(-4))
	whichOS()
	whenissat()
	tryDefer()

}

func test() {
	k := 3
	l := "wo cao"
	var x, y int = 3, 4
	var f float64 = math.Sqrt(float64(x*x + y*y))
	var z uint = uint(f)
	fmt.Println("Welcome to the playground!")
	fmt.Println("The time is", time.Now())
	fmt.Println(rand.Intn(10))
	fmt.Println(math.Pi)
	fmt.Println(math.Sqrt(9))
	fmt.Println(add(9, 12))
	fmt.Println(split(17))
	fmt.Println(java, python, ios, i, j, k, l, MaxInt, z, f)
}

func add(x int, y int) int {
	return x + y
}

func swap(x, y string) (string, string) {
	return y, x
}

func split(sum int) (x, y int) {

	x = sum * 4 / 9
	y = sum - x
	return

}
func flowFor() {

	sum := 1
	for sum < 1000 {
		sum += sum
		fmt.Println(sum)
	}
	//fmt.Println(sum)

}

func sqrt(x float64) string {
	if x < 0 {
		return sqrt(-x) + "i"
	}
	return fmt.Sprint(math.Sqrt(x))
}

func whichOS() {

	fmt.Print("Go runs on ")
	switch os := runtime.GOOS; os {
	case "darwin":
		fmt.Println("OS X.")
	case "linux":
		fmt.Println("Linux.")
	default:
		// freebsd, openbsd,
		// plan9, windows...
		fmt.Printf("%s.", os)
	}
	fmt.Println(runtime.GOROOT())

}

func whenissat() {
	fmt.Println("When is saturday?")
	today := time.Now().Weekday()
	switch time.Saturday {
	case today + 0:
		fmt.Println("today")
	case today + 1:
		fmt.Println("tomorrow")
	case today + 2:
		fmt.Println("in tow days")
	default:
		fmt.Println("far away")

	}

}
func tryDefer() {
	fmt.Println("counting")

	for i := 0; i < 10; i++ {
		defer fmt.Println(i)
	}

	fmt.Println("done")

}
