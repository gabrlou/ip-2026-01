package main

import f "fmt"

func main() {
	var (
		numeros [10]float64
	)

	for i := 0; i < len(numeros); i++ {
		f.Printf("Informe o %dº número: ", i+1)
		f.Scan(&numeros[i])
	}

	f.Print("Números na ordem inversa apresentada: ")

	for i := 9; i >= 0; i = i - 1 {
		if i == 0 {
			f.Print(numeros[i], ". ")
		} else {
			f.Print(numeros[i], ", ")
		}
	}
}

