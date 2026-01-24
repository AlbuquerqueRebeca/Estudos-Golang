package main

import "fmt"

func main() {

	idade := map[string]int{}
	idade["Rebeca"] = 35
	idade["Olavo"] = 9
	//fmt.Println(idade) // acessando a chave e valor
	//fmt.Println(idade["Rebeca"], 35)  //acessandoo valor pela chave
	//fmt.Println(idade["Olavo"], 9)

	anoNascimento := map[string]int{
		"Rebeca": 1990,
		"Olavo":  2017,
	}
	//fmt.Println(anoNascimento)
	fmt.Println(anoNascimento["Rebeca"])
	fmt.Println(anoNascimento["Olavo"])
}
