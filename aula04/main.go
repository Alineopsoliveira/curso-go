package main

import "fmt"

const a = "Hello, World!"

type ID int

var (
	b bool    = true
	c int     = 10
	d string  = "Aline"
	e float64 = 3.14
	f ID      = 1
)

func main() {

	fmt.Printf("O tipo é %T", e)
}
