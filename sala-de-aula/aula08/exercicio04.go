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
	valor_base := numero
	calculo := numero
	for i := numero; i > 1; i-- {
		calculo *= (valor_base - 1)
		valor_base -= 1
	}
	return calculo
}
