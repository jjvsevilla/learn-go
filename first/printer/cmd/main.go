package main

import "github.com/jjvsevilla/learn-go/first/printer"

func main() {
	printer.Hello()

	var version = printer.Version()
	printer.Print("Current Go version " + version)

}
