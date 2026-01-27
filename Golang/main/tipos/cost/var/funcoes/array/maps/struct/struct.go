package main

import "fmt"

func main() {

	type pessoa struct {
		nome  string
		idade int
	}

	//fmt.Println(pessoa{"Rebeca, 25"})
	//fmt.Println(pessoa{nome: "Rebeca", idade: 25})
	//fmt.Println(pessoa{nome: "Rebeca"})
	//fmt.Println(pessoa{idade: 25})

	p1 := pessoa{nome: "Rebeca", idade: 25}
	//fmt.Println(p1.nome)
	//fmt.Println(p1.idade)
	fmt.Println(p1.nome, p1.idade)

	p2 := pessoa{"Olavo", 8}
	//fmt.Println(p2.nome)
	//fmt.Println(p2.idade)
	fmt.Println(p2.nome, p2.idade)


	pessoas := []pessoa{}
	pessoas = append(pessoas, p1, p2)   //adicionando elementos 


   alunos := map[string] pessoa{
	alunos["programação"] = pessoas
	fmt.Println(alunos)
   }

}
