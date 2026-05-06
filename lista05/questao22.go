package main

import f "fmt"

func main() {

	codigos := make([]int, 10)
	saldos := make([]float64, 10)

	var codigo int
	var valor float64

	for i := 0; i < 10; i++ {

		for {
			f.Printf("Digite o código da %dª conta: ", i+1)
			f.Scan(&codigo)

			existe := false

			for j := 0; j < i; j++ {
				if codigos[j] == codigo {
					existe = true
				}
			}

			if existe == false {
				codigos[i] = codigo
				break
			} else {
				f.Print("Código já existe. Digite outro.\n")
			}
		}

		f.Printf("Digite o saldo da conta %d: ", codigos[i])
		f.Scan(&saldos[i])
	}

	op := 0

	for op != 4 {

		f.Print("\n1. Efetuar Depósito\n")
		f.Print("2. Efetuar Saque\n")
		f.Print("3. Consultar o ativo bancário\n")
		f.Print("4. Finalizar o programa.\n")
		f.Print("Escolha: ")

		f.Scan(&op)

		if op == 1 {
			f.Print("Digite o Código da conta: ")
			f.Scan(&codigo)

			pos := -1

			for i := 0; i < 10; i++ {
				if codigos[i] == codigo {
					pos = i
				}
			}

			if pos == -1 {
				f.Print("Conta não encontrada.\n")
			} else {
				f.Print("Digite o Valor do depósito: ")
				f.Scan(&valor)

				saldos[pos] = saldos[pos] + valor
				f.Print("Depósito realizado.\n")
			}

		} else if op == 2 {
			f.Print("Digite o Código da conta: ")
			f.Scan(&codigo)

			pos := -1

			for i := 0; i < 10; i++ {
				if codigos[i] == codigo {
					pos = i
				}
			}

			if pos == -1 {
				f.Print("Conta não encontrada.\n")
			} else {
				f.Print("Digite o Valor do saque: ")
				f.Scan(&valor)

				if saldos[pos] >= valor {
					saldos[pos] = saldos[pos] - valor
					f.Print("Saque realizado.\n")
				} else {
					f.Print("Saldo insuficiente.\n")
				}
			}

		} else if op == 3 {
			total := 0.0

			for i := 0; i < 10; i++ {
				total = total + saldos[i]
			}

			f.Printf("Ativo bancário: %.2f\n", total)

		} else if op != 4 {
			f.Print("Opção inválida.\n")
		}
	}

	f.Print("Programa encerrado.\n\n")
}
