package main

import "fmt"

func main() {
	mult := soma(40, 13)
	fmt.Println(mult)

	sub := subtracao(40, 20)
	fmt.Println(sub)

	div := divisao(50, 23)
	fmt.Println(div)

}

func soma(x int, y int) int {
	return x + y
}

func subtracao(x int, y int) int {
	return x - y
}

func divisao(x int, y int) int {
	return x / y
}
