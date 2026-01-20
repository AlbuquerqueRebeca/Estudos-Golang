package main

import "fmt"

func main() {
	fmt.Printf("Type: %T - value: %v\n", true, true)   //tipo booleano
	fmt.Printf("Type: %T - value: %v\n", false, false) //tipo booleano
	fmt.Printf("Type: %T - value: %v\n", 1, 1)         //tipo inteiro
	fmt.Printf("Type: %T - value: %v\n", 1.234, 1.234) //tipo float
}
