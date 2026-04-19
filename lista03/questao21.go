package main

import f "fmt"

func main() {
	var b, n int

	for {
		f.Print("\nDigite o valor da base b: ")
		f.Scan(&b)
		if b < 2 {
			f.Print("\nNúmero inválido. Digite um número maior ou igual a 2 para a base.\n")
			continue
		} else {
			break
		}
	}

	for {
		f.Print("Digite o valor do expoente n: ")
		f.Scan(&n)
		if n <= 1 {
			f.Print("\nNúmero inválido. Digite um número maior que 1 para o expoente.\n")
			continue
		} else {
			break
		}
	}

	resultado := potencia(b, n)
	f.Printf("\n%d ^ %d = %d\n\n", b, n, resultado)
}

func potencia(b, n int) int {
	if n == 1 {
		return b
	}
	return b * potencia(b, n-1)
}
