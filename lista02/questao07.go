package main

import f "fmt"

func main() {
	var A, B, C, temp, MENOR, INTER, MAIOR int

	f.Print("Digite o valor de A: ")
	f.Scanln(&A)

	f.Print("Digite o valor de B: ")
	f.Scanln(&B)

	f.Print("Digite o valor de C: ")
	f.Scanln(&C)

	if A > B {
		temp = A
		A = B
		B = temp
	}
	if B > C {
		temp = B
		B = C
		C = temp
	}
	if A > B {
		temp = A
		A = B
		B = temp
	}

	MENOR = A
	INTER = B
	MAIOR = C

	f.Print("\nMenor: ", MENOR, "\nIntermediário: ", INTER, "\nMaior: ", MAIOR)
}
