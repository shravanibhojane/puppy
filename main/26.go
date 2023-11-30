package main

import "fmt"

var y int

func main() {
	fmt.Println(y)

	z := 42
	fmt.Println(z)

	a, b := 43, "happiness"
	fmt.Println(a, b)
	var c float32 = 42.42
	fmt.Println(c)

	fmt.Printf("%v \t %T", c, c)
}
