package main

import f "fmt"

func main() {

	idades := make([]int, 50)

	f.Print("\n")
	for i := 0; i < 50; i++ {
		f.Printf("Digite a %dª idade: ", i+1)
		f.Scan(&idades[i])
	}

	moda := 0
	qtdVezes := 0

	for i := 0; i < 50; i++ {
		contador := 0
		jaContado := false

		for k := 0; k < i; k++ {
			if idades[k] == idades[i] {
				jaContado = true
				break
			}
		}

		if jaContado {
			continue
		}

		for j := 0; j < 50; j++ {
			if idades[j] == idades[i] {
				contador++
			}
		}

		if contador > qtdVezes {
			qtdVezes = contador
			moda = idades[i]
		}
	}

	f.Printf("\nModa das idades: %d (aparece %d vezes)\n\n", moda, qtdVezes)
}
