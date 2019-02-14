package printer

import "fmt"
import "runtime"

// Hello is an exported function
func Hello() {
	fmt.Println("exported hello")
}

// Print function print any text
func Print(text string) {
	fmt.Println(text)
}

// Version function return the current Go version
func Version() string {
	return runtime.Version()
}
