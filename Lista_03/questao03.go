package main

import f "fmt"

func main() {
	var (
		salario_carlos_inicial, salario_joao_inicial, salario_carlos_final, salario_joao_final float64
		meses                                                                                  int
	)

	f.Print("\nDigite o salário do funcionário Carlos: ")
	f.Scan(&salario_carlos_inicial)

	salario_joao_inicial = salario_carlos_inicial / 3

	meses = 0
	salario_carlos_final = salario_carlos_inicial
	salario_joao_final = salario_joao_inicial
	for salario_joao_final < salario_carlos_final {
		salario_carlos_final *= 1.02
		salario_joao_final *= 1.05
		meses++
	}
	f.Printf("\nSalário de Carlos Inicial: R$ %.2f | Salário de João Inicial: R$ %.2f", salario_carlos_inicial, salario_joao_inicial)
	f.Printf("\nSalário de Carlos Final: R$ %.2f | Salário de João Final: R$ %.2f", salario_carlos_final, salario_joao_final)
	f.Printf("\nTotal de meses: %d\n\n", meses)
}
