package main

import "fmt"

func main() {
	var minhaVar interface{} = "Aline Oliveira"

	println(minhaVar.(string))

	res, ok := minhaVar.(int)

	fmt.Printf("O valor da variável é %v e o resultado é %v\n", res, ok)
	res2 := minhaVar.(int)
	fmt.Printf("O valor da variável é %v\n", res2)
}
