package main

import "fmt"

func main() {
	fmt.Println(Soma(10, 10))
}

// Soma retorna a soma dos números
func Soma(a, b int) int {
	return a + b
}
