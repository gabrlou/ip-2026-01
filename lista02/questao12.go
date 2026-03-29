package main

import f "fmt"

func main() {
	var idade int
	f.Print("Digite sua idade: ")
	f.Scan(&idade)

	if idade < 0 {
		f.Println("Idade inválida")
	} else if idade >= 0 && idade <= 2 {
		f.Println("Recém-nascido")
	} else if idade >= 3 && idade <= 11 {
		f.Println("Criança")
	} else if idade >= 12 && idade <= 19 {
		f.Println("Adolescente")
	} else if idade >= 20 && idade <= 55 {
		f.Println("Adulto")
	} else {
		f.Println("Idoso")
	}
}
