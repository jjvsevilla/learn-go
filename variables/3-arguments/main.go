package main

import "fmt"
import "os"

func main() {
	// fmt.Println("arguments")
	// fmt.Printf("%#v\n", os.Args)
	// fmt.Println("path ", os.Args[0])
	// fmt.Println("1st argument ", os.Args[1])
	// fmt.Println("number of params ", len(os.Args))

	name := os.Args[1]
	fmt.Println("hello great", name, "!")

	name, age := "gandalf", 2019
	fmt.Println("my name is", name)
	fmt.Println("my age is", age)
	fmt.Println("btw, you shall pass!")
}
