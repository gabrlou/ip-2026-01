package main

import f "fmt"

func main() {
	var A, B, C, delta float64

	//Entrada
	f.Print("Digite o coeficiente \"A\": ")
	f.Scanln(&A)

	f.Print("Digite o coeficiente \"B\": ")
	f.Scanln(&B)

	f.Print("Digite o coeficiente \"C\": ")
	f.Scanln(&C)

	//Cálculo de Delta
	delta = (B * B) - 4*A*C

	//Saída
	f.Print("\n")
	f.Print("O VALOR DE DELTA E = ", f.Sprintf("%.2f", delta))
}
