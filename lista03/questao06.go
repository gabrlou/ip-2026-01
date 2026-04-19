package main

import f "fmt"

func main() {
	var (
		num       int
		resultado = false
	)

	for {
		f.Print("\n	Digite um número inteiro positivo: ")
		f.Scan(&num)
		if num < 0 {
			f.Print("\nNúmero inválido! Digite um valor maior que zero.")
			continue
		} else {
			break
		}
	}

	for i := 1; (i * (i + 1) * (i + 2)) <= num; i++ {
		if i*(i+1)*(i+2) == num {
			f.Printf("	O número é triangular pois %d * %d * %d = %d\n\n", i, i+1, i+2, num)
			resultado = true
			break
		} else {
			continue
		}
	}
	if resultado == false {
		f.Print("	O número informado não é triangular.\n\n")
	}
}
