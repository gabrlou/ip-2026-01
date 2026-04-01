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
   
   result := MEDIA(n1,n2,n3)
   f.Printf("\nA média entre os números %d, %d e %d é: %.2f\n", n1, n2, n3, result)
}

func MEDIA(X, Y, Z int) float64 {
	media := float64(X + Y + Z) / 3
	return media
}
