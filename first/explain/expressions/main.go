package main

import (
	"fmt"
	"runtime"
)

func main() {
	fmt.Println("hola!")
	fmt.Println(runtime.NumCPU() + 1)
}
