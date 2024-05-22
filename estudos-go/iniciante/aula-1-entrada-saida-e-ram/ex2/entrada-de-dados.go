package main

import "fmt"

func main() {
	var nome string // Declara a variavel nome
	fmt.Println("Digite seu nome: ")
	fmt.Scanf("%s", &nome)
	//var sala string
	//sala = "313"
	sala := "313" // o comando := declara e atribui valor a uma variavel
	fmt.Printf("Seja bem vindo(a), %s! Sua sala é %s!", nome, sala)
}
