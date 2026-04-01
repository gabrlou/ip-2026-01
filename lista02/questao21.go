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

	if media_final < 4 {
		f.Printf("\nIdentificação do aluno: %d\nNota 1: %.2f\nNota 2: %.2f\nNota 3: %.2f\nMedia obtida nos exercícios: %.2f\nMédia de aproveitamento: %.2f\nConceito: E\nREPROVADO\n", id, n1, n2, n3, media_ex, media_final)
	} else if media_final >= 4 && media_final < 6 {
		f.Printf("\nIdentificação do aluno: %d\nNota 1: %.2f\nNota 2: %.2f\nNota 3: %.2f\nMedia obtida nos exercícios: %.2f\nMédia de aproveitamento: %.2f\nConceito: D\nREPROVADO\n", id, n1, n2, n3, media_ex, media_final)
	} else if media_final >= 6 && media_final < 7.5 {
		f.Printf("\nIdentificação do aluno: %d\nNota 1: %.2f\nNota 2: %.2f\nNota 3: %.2f\nMedia obtida nos exercícios: %.2f\nMédia de aproveitamento: %.2f\nConceito: C\nAPROVADO\n", id, n1, n2, n3, media_ex, media_final)
	} else if media_final >= 7.5 && media_final < 9 {
		f.Printf("\nIdentificação do aluno: %d\nNota 1: %.2f\nNota 2: %.2f\nNota 3: %.2f\nMedia obtida nos exercícios: %.2f\nMédia de aproveitamento: %.2f\nConceito: B\nAPROVADO\n", id, n1, n2, n3, media_ex, media_final)
	} else {
		f.Printf("\nIdentificação do aluno: %d\nNota 1: %.2f\nNota 2: %.2f\nNota 3: %.2f\nMedia obtida nos exercícios: %.2f\nMédia de aproveitamento: %.2f\nConceito: A\nAPROVADO\n", id, n1, n2, n3, media_ex, media_final)
	}
}
