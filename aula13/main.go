package main

import "fmt"

type Cliente struct {
	Nome string
}

func (c Cliente) andou() {
	c.Nome = "Aline Oliveira"
	fmt.Printf("A cliente %v andou\n", c.Nome)
}
func main() {
	aline := Cliente{
		Nome: "aline",
	}
	aline.andou()
	fmt.Printf("O valor da struct com nome %v \n", aline.Nome)
}
