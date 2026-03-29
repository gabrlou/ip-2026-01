package main

import f "fmt"

func main() {
	var numero, i int

	//Entrada
	f.Print("Digite um número inteiro maior que 5 e menor que 2000: ")
	f.Scanln(&numero)

	if numero <= 5 || numero >= 2000 {
		f.Print("\nNÚMERO INVÁLIDO!\nDigite um número maior que 5 e menor que 2000.\n")
		return
	}

	f.Print("\n")

	//Processamento e Saída
	for i = 2; i <= numero; i += 2 {
		f.Print(i, "^2 = ", i*i, "\n")
	}
}
