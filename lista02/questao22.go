package main

import f "fmt"

func main() {
	var (
		salario_bruto, salario_liq, salario_extra, descont_inss, imposto_renda, horas_extras float64
		matricula                                                                            int
	)
	const (
		SALARIO_MIN      int = 788
		VALOR_HORA_EXTRA int = 10
	)

	f.Print("Digite o número de matrícula do funcionário: ")
	f.Scanln(&matricula)

	f.Print("Digite a quantidade de horas extras trabalhadas: ")
	f.Scanln(&horas_extras)

	salario_extra = horas_extras * float64(VALOR_HORA_EXTRA)
	salario_bruto = 3*float64(SALARIO_MIN) + float64(salario_extra)

	if salario_bruto > 1500 {
		descont_inss = salario_bruto * 0.12
	}
	if salario_bruto > 2000 {
		imposto_renda = salario_bruto * 0.2
	}

	salario_liq = salario_bruto - descont_inss - imposto_renda

	f.Printf("\nMatrícula: %d\nSalário bruto: R$ %.2f\nSalário líquido: R$ %.2f\n", matricula, salario_bruto, salario_liq)
}
