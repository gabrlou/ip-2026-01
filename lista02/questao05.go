package main
import f "fmt"

func main() {
    
    var numero int
    
    f.Println("Digite um número do tipo inteiro: ")
    f.Scan(&numero)
    f.Println("\n")
    
    if numero % 5 == 0 {
        f.Println("O número é divisível por 5.")
    } else {
        f.Println("O número não é divisível por 5.")
    }
}
