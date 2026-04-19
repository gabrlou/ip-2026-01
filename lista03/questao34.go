package main

import f "fmt"

func main() {
	var n1, n2 int

	for {
		f.Print("\nDigite o valor de N1: ")
		f.Scan(&n1)

		if n1 <= 0 {
			f.Println("Número inválido. Digite um número maior que zero.")
			continue
		} else {
			break
		}
	}

	for {
		f.Print("Digite o valor de N2: ")
		f.Scan(&n2)

		if n2 <= 0 {
			f.Println("Número inválido. Digite um número maior que zero.")
			continue
		} else {
			break
		}
	}

	resultado := mmc(n1, n2)
	f.Printf("\nMMC = %d\n\n", resultado)
}

// Função para calcular o MDC usando o algoritmo de Euclides
func mdc(n1, n2 int) int {
	for n2 != 0 {
		n1, n2 = n2, n1%n2
	}
	return n1
}

// Função para calcular o MMC
func mmc(n1, n2 int) int {
	return (n1 * n2) / mdc(n1, n2)
}
