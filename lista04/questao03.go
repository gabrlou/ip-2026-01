package main

import f "fmt"

func main() {

	var quantidade int
	f.Print("\nQuantos números serão inseridos? ")
	f.Scan(&quantidade)

	numeros := make([]int, quantidade)

	for i := 0; i < quantidade; i++ {
		f.Printf("Digite o %dº número: ", i+1)
		f.Scan(&numeros[i])
	}

	inverterArray(numeros)

	f.Println("\nArray invertido:")

	for i := 0; i < quantidade; i++ {
		f.Print(numeros[i], " ")
	}
}

func inverterArray(numeros []int) {
	if len(numeros) <= 1 {
		return
	}
	numeros[0] = numeros[len(numeros)-1]
	numeros[len(numeros)-1] = numeros[0]

	inverterArray(numeros[1 : len(numeros)-1])
}
