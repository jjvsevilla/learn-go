package main

import (
	"fmt"
	"os"
	"strconv"
)

func main() {
	// var (
	// 	myAge   = 30
	// 	yourAge = 35
	// 	average float64
	// )
	// average = float64(myAge+yourAge) / 2
	// fmt.Println(average)

	// ratio := 1.0 / 10
	// // fmt.Printf("%.60f", ratio)
	// fmt.Println(ratio)

	// // precedence
	// var m = 2 + ((2 * 4) / 2)
	// fmt.Println(m)

	// celsius := 35.
	// fahrenheit := (9*celsius + 160) / 5

	// fmt.Printf("%g C is %g F\n", celsius, fahrenheit)

	// n := 10
	// n = n + 1
	// n += 1
	// n++
	// fmt.Println(n)

	arg := os.Args[1]
	feet, _ := strconv.ParseFloat(arg, 64)
	meters := feet * 0.3048
	// fmt.Printf("%f feet is %f meters.\n", feet, meters)
	// %g remove unnecesaried fractional parts
	fmt.Printf("%g feet is %g meters.\n", feet, meters)
}
