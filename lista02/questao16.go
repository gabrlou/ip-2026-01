package main

import (
	f "fmt"
	m "math"
)

func main() {
	var A, B, C, delta, raiz_1, raiz_2, parte_a, parte_b float64

	f.Print("Digite o valor de A:")
	f.Scanln(&A)

	f.Print("Digite o valor de B:")
	f.Scanln(&B)

	f.Print("Digite o valor de C:")
	f.Scanln(&C)

	delta = (B * B) - (4 * A * C)
	if A == 0 { //Verificar se é uma equação do 2º Grau
		f.Println("\nValores inválidos!\nNão é uma equação do 2º Grau.")
		return
	} else if delta < 0 { //RAIZES IMAGINARIAS
		parte_a = (-B) / (2 * A)
		parte_b = (m.Sqrt(-delta)) / (2 * A)

		if parte_a == 0 {
			f.Printf("\nValor de x1: + %.2fi\nValor de x2: - %.2fi", parte_b, parte_b)
			f.Println("\nClassificação: RAÍZES IMAGINÁRIAS")
		} else {
			f.Printf("\nValor de x1: %.2f + %.2fi\nValor de x2: %.2f - %.2fi", parte_a, parte_b, parte_a, parte_b)
			f.Println("\nClassificação: RAÍZES IMAGINÁRIAS")
		}
	} else if delta == 0 { //RAIZ UNICA
		raiz_1 = (-B) / (2 * A)

		f.Printf("\nValor de x1 e x2: %.2f", raiz_1)
		f.Println("\nClassificação: RAÍZ ÚNICA")
	} else { //RAIZES DISTINTAS
		raiz_1 = (-B + m.Sqrt(delta)) / (2 * A)
		raiz_2 = (-B - m.Sqrt(delta)) / (2 * A)

		f.Printf("\nValor de x1: %.2f\nValor de x2: %.2f", raiz_1, raiz_2)
		f.Println("\nClassificação: RAÍZES DISTINTAS")
	}
}
