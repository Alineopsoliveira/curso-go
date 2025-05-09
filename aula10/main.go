package main

import "fmt"

func main() {

	fmt.Println(sum(1, 85, 94, 67, 7, 2, 13))

}

func sum(numeros ...int) int {
	total := 0
	for _, numero := range numeros {
		total += numero
	}
	return total
}
