package main

import f "fmt"

func main() {

	var (
		id, peso                     int
		maior_peso, menor_peso       int
		id_maior_peso, id_menor_peso int
	)

	for i := 0; i < 90; i++ {

		f.Printf("Boi nº %d - Digite o ID: ", i+1)
		f.Scan(&id)

		for {
			f.Printf("Boi nº %d - Digite o peso: ", i+1)
			f.Scan(&peso)

			if peso <= 0 {
				f.Println("Peso inválido. Digite um valor maior que zero.\n")
				continue
			} else {
				break
			}
		}

		if i == 0 {
			maior_peso = peso
			menor_peso = peso
			id_maior_peso = id
			id_menor_peso = id
		} else {
			// maior peso
			if peso > maior_peso {
				maior_peso = peso
				id_maior_peso = id
			}

			// menor peso
			if peso < menor_peso {
				menor_peso = peso
				id_menor_peso = id
			}
		}
	}

	f.Println("\n	RESULTADO FINAL:")
	f.Printf("	Boi mais gordo: ID %d com %d kg\n", id_maior_peso, maior_peso)
	f.Printf("	Boi mais magro: ID %d com %d kg\n", id_menor_peso, menor_peso)
}
