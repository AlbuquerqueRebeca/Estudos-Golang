package main

import "fmt"

func main() {
	var array [3]string
	array[0] = "Rebeca"
	array[1] = "Guanaes"
	array[2] = "Albuquerque"
	//fmt.Println(array[0], array[1], array[2])
	//fmt.Println(array[2])

	//sequencia de inteiros
	//numInteiro := [6]int{1, 2, 3, 4, 5, 6}
	//fmt.Println(numInteiro[0])
	//fmt.Println(numInteiro[1])
	//fmt.Println(numInteiro[2])
	//fmt.Println(numInteiro[3])
	//fmt.Println(numInteiro[4])
	//fmt.Println(numInteiro[5])

	//var slice [] string
	slice := make([]string, 4) //dizendo um tamanho aproximado
	slice[0] = "Rebeca"
	slice[1] = "Guanaes"
	fmt.Println(slice[0])
	fmt.Println(slice[1])

}
