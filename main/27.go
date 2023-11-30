package main

import "fmt"

func main() {
	s, i, f := "happiness", 42, 42.42
	fmt.Printf("%v is of type %T", s, s)
	fmt.Printf("%v is of type %T", i, i)
	fmt.Printf("%v is of type %T", f, f)
}
