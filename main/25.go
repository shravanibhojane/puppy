package main

import "fmt"

type bytesize int

const (
	_           = iota
	KB bytesize = 1 << (10 * iota)
	MB
	GB
	TB
	PB
)

func main() {
	fmt.Printf("%d \t\t\t %b\n", KB, KB)
	fmt.Printf("%d \t\t\t %b\n", MB, MB)
	fmt.Printf("%d \t\t\t %b\n", TB, TB)
	fmt.Printf("%d \t\t\t %b\n", PB, PB)

}
