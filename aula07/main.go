package main

import "fmt"

func main() {
	salarios := map[string]int{"João": 1000, "Maria": 2000, "Pedro": 1500}

	delete(salarios, "João")
	salarios["Joca"] = 3000

	for nome, salario := range salarios {
		fmt.Printf("O salario de %s é %d\n", nome, salario)
	}

	for _, salario := range salarios {
		fmt.Printf("O salario é %d\n", salario)
	}
}
