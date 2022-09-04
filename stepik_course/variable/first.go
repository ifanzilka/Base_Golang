package main

import "fmt"

func tmp() {
	a := 5
	fmt.Println(a)
}

func main() {
	var hello string
	var x int
	var y float64
	var max, min int

	var x int = 10
	var c string = "Hello World!"
	var z float64 = 1.045

	var (
		name string = "Dima"
		age  int    = 23
	)

	var m float32 = 10.0 / 6

	fmt.Println(m)
	tmp()
}
