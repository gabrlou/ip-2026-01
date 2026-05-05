package main

import f "fmt"

func main() {

	A := make([]int, 10)

	f.Print("\n")
	for i := 0; i < 10; i++ {
		f.Printf("Digite o %dº número: ", i+1)
		f.Scan(&A[i])
	}

	f.Print("\nElementos repetidos:\n")

	existe := false

	for i := 0; i < 10; i++ {
		contador := 0
		jaContado := false

		for j := 0; j < i; j++ {
			if A[j] == A[i] {
				jaContado = true
				break
			}
		}

		if jaContado {
			continue
		}

		for k := 0; k < 10; k++ {
			if A[k] == A[i] {
				contador++
			}
		}

		if contador > 1 {
			f.Printf("Número %d aparece %d vezes\n", A[i], contador)
			existe = true
		}
	}

	if existe == false {
		f.Print("Nenhum número repetido.\n")
	}

	f.Print("\n")
}
