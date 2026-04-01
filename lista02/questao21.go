package main
import f "fmt"

func main() {
	var (
		n1, n2, n3, media_ex, media_final float64
		id int
	)

	f.Print("Digite o número de identificação do aluno: ")
	f.Scanln(&id)
	
	f.Print("Digite a 1ª nota do aluno: ")
	f.Scanln(&n1)

	f.Print("Digite a 2ª nota do aluno: ")
	f.Scanln(&n2)

	f.Print("Digite a 3ª nota do aluno: ")
	f.Scanln(&n3)

	f.Print("Digite a média obtida nos exercícios: ")
	f.Scanln(&media_ex)

	media_final = (n1 + (n2 * 2) + (n3 * 3) + media_ex) / 7
	f.Printf("\nIdentificação do aluno: %d\nMédia de aproveitamento: %.2f", id, media_final)
}
