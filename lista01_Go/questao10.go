package main

import f "fmt"

func main() {
	var a, b, c, d, determinante float64

	//Entrada
	f.Print("Digite o elemente \"a\" da matriz 2x2: ")
	f.Scanln(&a)

	f.Print("Digite o elemento \"b\" da matriz 2x2: ")
	f.Scanln(&b)

	f.Print("Digite o elemento \"c\" da matriz 2x2: ")
	f.Scanln(&c)

	f.Print("Digite o elemento \"d\" da matriz 2x2: ")
	f.Scanln(&d)

	//Cálculo de Determinante
	determinante = (a * d) - (b * c)

	//Saída
	f.Print("\n")
	f.Print("O VALOR DO DETERMINANTE E = ", f.Sprintf("%.2f", determinante))
}
