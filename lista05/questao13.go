package main

import f "fmt"

func main() {

	var numero, meses int

	qtdMeses1, qtdMeses2, qtdMeses3 := 999999, 999999, 999999
	num1, num2, num3 := 0, 0, 0

	f.Print("\nDigite (numero do empregado, meses de trabalho). Digite 0 0 para sair.\n")

	for {
		f.Print(">> ")
		f.Scan(&numero, &meses)

		if numero == 0 && meses == 0 {
			break
		}

		if meses < qtdMeses1 {
			qtdMeses3, num3 = qtdMeses2, num2
			qtdMeses2, num2 = qtdMeses1, num1
			qtdMeses1, num1 = meses, numero

		} else if meses < qtdMeses2 {
			qtdMeses3, num3 = qtdMeses2, num2
			qtdMeses2, num2 = meses, numero

		} else if meses < qtdMeses3 {
			qtdMeses3, num3 = meses, numero
		}
	}

	f.Print("\n3 empregados mais recentes:\n")
	f.Printf("1º -> Empregado %d (%d meses)\n", num1, qtdMeses1)
	f.Printf("2º -> Empregado %d (%d meses)\n", num2, qtdMeses2)
	f.Printf("3º -> Empregado %d (%d meses)\n", num3, qtdMeses3)

	f.Print("\n\n")
}
