package main

import f "fmt"

func main() {
	var salario, salario_reajustado float64

	//Entrada
	f.Print("Digite o salário do funcionário: ")
	f.Scanln(&salario)

	if salario <= 300 {
		salario_reajustado = salario + (salario * 0.5)
		f.Print("\nSALARIO COM REAJUSTE = ", f.Sprintf("%.2f", salario_reajustado), "\n")
	} else {
		salario_reajustado = salario + (salario * 0.3)
		f.Print("\nSALARIO COM REAJUSTE = ", f.Sprintf("%.2f", salario_reajustado), "\n")
	}
}
