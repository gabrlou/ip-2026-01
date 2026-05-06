package main

import f "fmt"

func main() {

	janela := make([]int, 24)
	corredor := make([]int, 24)

	op := 1

	for op != 0 {

		f.Print("\n1 - Janela\n")
		f.Print("2 - Corredor\n")
		f.Print("0 - Sair\n")
		f.Print("Escolha: ")
		f.Scan(&op)

		if op == 1 {

			livre := false

			f.Print("Poltronas disponíveis (Janela): ")

			for i := 0; i < 24; i++ {
				if janela[i] == 0 {
					f.Printf("%d ", i)
					livre = true
				}
			}

			if livre == false {
				f.Print("\nNão há poltronas livres na janela.\n")
			} else {
				var escolha int
				f.Print("\nEscolha a poltrona: ")
				f.Scan(&escolha)

				if escolha >= 0 && escolha < 24 {
					if janela[escolha] == 0 {
						janela[escolha] = 1
						f.Print("Poltrona reservada.\n")
					} else {
						f.Print("Poltrona já ocupada.\n")
					}
				} else {
					f.Print("Posição inválida.\n")
				}
			}

		} else if op == 2 {

			livre := false

			f.Print("Poltronas disponíveis (Corredor): ")

			for i := 0; i < 24; i++ {
				if corredor[i] == 0 {
					f.Printf("%d ", i)
					livre = true
				}
			}

			if livre == false {
				f.Print("\nNão há poltronas livres no corredor.\n")
			} else {
				var escolha int
				f.Print("\nEscolha a poltrona: ")
				f.Scan(&escolha)

				if escolha >= 0 && escolha < 24 {
					if corredor[escolha] == 0 {
						corredor[escolha] = 1
						f.Print("Poltrona reservada.\n")
					} else {
						f.Print("Poltrona já ocupada.\n")
					}
				} else {
					f.Print("Posição inválida.\n")
				}
			}

		} else if op != 0 {
			f.Print("Opção inválida.\n")
		}

		total := 0
		for i := 0; i < 24; i++ {
			total = total + janela[i]
			total = total + corredor[i]
		}

		if total == 48 {
			f.Print("\nÔnibus completamente lotado.\n")
			op = 0
		}
	}

	f.Print("\nPrograma encerrado.\n")
}
