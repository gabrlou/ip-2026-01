package main
import f  "fmt"
func main() {
   
   n1:= 0
   f.Print("Digite um número inteiro: ")
   f.Scan(&n1)
   
   result := Fatorial(n1)
   f.Printf("\nNúmero: %d\nFatorial: %d\n", n1, result)
}

func Fatorial(numero int) int {
	if numero == 1 {
		return numero
	}
	return numero * Fatorial(numero - 1)
}
