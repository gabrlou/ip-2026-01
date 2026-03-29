package main

import f "fmt"

func main() {
	var numero int

	//Entrada
	f.Print("Digite um número inteiro: ")
	f.Scanln(&numero)

	//Processamento e Saída
	if numero%3 == 0 && numero%5 == 0 {
		f.Print("\nO NUMERO E DIVISIVEL\n")
	} else {
		f.Print("\nO NUMERO NAO E DIVISIVEL\n")
	}
}
