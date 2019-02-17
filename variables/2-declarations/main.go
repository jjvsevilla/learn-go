package main

import "fmt"

func main() {
	/*
		var (
			speed     int
			heat      float64
			off       bool
			brand     string
			unusedVar int
		)
		fmt.Println(off)
		fmt.Println(heat)
		fmt.Println(speed)
		fmt.Printf("%q\n", brand)
		var _ = unusedVar // blank-identifier

		var hours, minutes int
		fmt.Println(hours, minutes)

		salary, employed := 5000, true
		fmt.Println(salary, employed)

		name, lastname := "nicola", "tesla"
		fmt.Println(name, lastname)

		birth, deatg := 1856, 1943
		fmt.Println(birth, deatg)

		on, off := true, false
		fmt.Println(on, off)

		degree, ratio, heat := 10.55, 30.5, 20.
		fmt.Println(degree, ratio, heat)

		nFiles, valid, msg := 10, true, "hello"
		fmt.Println(nFiles, valid, msg)

		var hasNintendo bool
		hasNintendo, isPlaying := true, false
		fmt.Println(hasNintendo, isPlaying)

		name, age := "juanjo", 33
		fmt.Println(name, age)

		name, birthh := "albert", 1879
		fmt.Println(name, birthh)
	*/

	/*
		var (
			name   string
			age    int
			famous bool
		)
		name = "newton"
		age = 84
		famous = true
		fmt.Println(name, age, famous)

		var prevName string
		prevName = name
		name = "juanjo"
		fmt.Println("prevname " + prevName)
		fmt.Println("name " + name)
	*/

	// var (
	// 	speed int
	// 	now   time.Time
	// )
	// speed, now = 100, time.Now()
	// fmt.Println(speed, now)

	// path package
	// var (
	// 	speed     = 100
	// 	prevSpeed = 50
	// )
	// speed, prevSpeed = prevSpeed, speed
	// fmt.Println(speed, prevSpeed)

	// var file string
	// _, file = path.Split("css/min/main.scss")
	// // fmt.Println("dir " + dir)
	// fmt.Println("file  " + file)

	// _, filee := path.Split("css/min/main.scss")
	// fmt.Println("filee " + filee)

	// type conversion
	var apple int
	var orange int32
	apple = int(orange)

	// isDelicious := bool(orange)
	orange = 65
	color := string(orange)
	fmt.Println(color)

	_ = apple
}
