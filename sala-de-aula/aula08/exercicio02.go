package main
import f  "fmt"
func main() {
   
   n1, n2, n3 := 0, 0, 0
   f.Print("Digite o primeiro número: ")
   f.Scan(&n1)
   f.Print("Digite o segundo número: ")
   f.Scan(&n2)
   f.Print("Digite o terceiro número: ")
   f.Scan(&n3)
   
   result := maior(n1,n2,n3)
   f.Printf("\nO maior número entre %d, %d e %d é: %d\n", n1, n2, n3, result)
}

func maior(numero1, numero2, numero3 int) int {
	maior_entre := 0
	if numero1 > numero2 && numero1 > numero3 {
		maior_entre = numero1
	} else if numero2 > numero1 && numero2 > numero3 {
		maior_entre = numero2
	} else {
		maior_entre = numero3
	}
	return maior_entre
}
