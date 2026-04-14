package main
import f "fmt"

func main() {
	var N int

	for {
		f.Print("\nDigite a quantidade N de termos da série [1 4 9 16 25 36 ...] a serem exibidos: ")
		f.Scan(&N)
		if N <= 0 {
			f.Print("\nNúmero inválido. Digite um valor maior que zero.\n")
			continue
		} else {
			break
		}
	}

	razao := 3
	termo := 1
	for i := 0; i < N; i++ {
		f.Printf("%d ", termo)

		termo += razao
		razao += 2
	}
	f.Print("\n\n")
}
