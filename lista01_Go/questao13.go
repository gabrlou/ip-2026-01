package main

import f "fmt"

func main() {
	var nota float64

	//Entrada
	f.Print("Digite a nota a ser convertida para conceito: ")
	f.Scanln(&nota)

	//Processamento e Saída
	if nota >= 0 && nota < 6 {
		f.Print("\nNOTA = ", nota, "   CONCEITO = D\n")
	} else if nota >= 6 && nota < 7.5 {
		f.Print("\nNOTA = ", nota, "   CONCEITO = C\n")
	} else if nota >= 7.5 && nota < 9 {
		f.Print("\nNOTA = ", nota, "   CONCEITO = B\n")
	} else {
		f.Print("\nNOTA = ", nota, "   CONCEITO = A\n")
	}
}
