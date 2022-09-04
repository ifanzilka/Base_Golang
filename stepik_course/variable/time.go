package main

import "fmt"

func main() {
	var a int
	var min, hour int

	fmt.Scan(&a)
	hour = a / 30
	min = 2 * (a % 30)

	fmt.Println("It is", hour, "hours", min, "minutes.")
}
