package main

import "fmt"

func main() {
	// format specifier
	// var brand = "Google"
	// fmt.Printf("%q\n", brand)

	// println vs printf
	// var ops, ok, fail int
	// ops, ok, fail = 2500, 2000, 500
	// fmt.Println(
	// 	"total:", ops, "success:", ok, "/", fail,
	// )

	// fmt.Printf(
	// 	"total: %d success: %d / %d\n",
	// 	ops, ok, fail,
	// )

	// escape sequences character \
	// fmt.Println("hi\nhow are \"juan\"?")

	// printing types with %T
	// var speed int
	// var heat float64
	// var off bool
	// var brand string

	// fmt.Printf("%T\n", speed)
	// fmt.Printf("%T\n", heat)
	// fmt.Printf("%T\n", off)
	// fmt.Printf("%T\n", brand)

	// verbs for each type
	var (
		planet   = "venus"
		distance = 261
		orbital  = 224.701
		hasLife  = false
	)
	fmt.Printf("Planet: %s\n", planet)
	fmt.Printf("Distance: %d millions kms\n", distance)
	fmt.Printf("Orbital period: %f days\n", orbital)
	fmt.Printf("Does %s have life? %t\n",
		planet, hasLife)

	// %v swiss-army-knife can print any type of value
	fmt.Printf(
		"%v is %v away. Think! %[2]v kms! %[1]v OMG!\n",
		planet, distance,
	)
	fmt.Printf("Orbital period: %f days\n", orbital)
	fmt.Printf("Orbital period: %.1f days\n", orbital)
	fmt.Printf("Orbital period: %.2f days\n", orbital)
	fmt.Printf("Orbital period: %.3f days\n", orbital)
	fmt.Printf("Orbital period: %.4f days\n", orbital)
}
