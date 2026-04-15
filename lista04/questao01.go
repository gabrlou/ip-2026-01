package main

import f "fmt"

func main() {
	var x, n int

	for {
		f.Print("\nDigite o valor da base x: ")
		f.Scan(&x)
		if x < 2 {
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

	resultado := potencia(x, n)
	f.Printf("\n%d ^ %d = %d\n\n", x, n, resultado)
}

func potencia (x, n int) int {
	if n == 0 {
		return 1
	}
	return x * potencia(x, n-1)
}
