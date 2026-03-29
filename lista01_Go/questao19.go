package main

import f "fmt"

func main() {
	var numero, contador int
	var soma float64

	//Entrada
	f.Print("Digite um número inteiro, positivo e maior que 1: ")
	f.Scanln(&numero)

	if numero <= 0 || numero == 1 {
		f.Print("\nNumero invalido!\n")
		return
	}

	soma = 0
	contador = 1

	//Processamento
	for contador <= numero {
		soma += 1 / float64(contador)
		contador += 1
	}

	//Saída
	f.Print("\n", f.Sprintf("%.6f", soma), "\n")
}
