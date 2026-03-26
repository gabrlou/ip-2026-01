package main
import f "fmt"

func main() {
    
    var a, b int
    
    f.Println("Digite um número inteiro \"A\": ")
    f.Scan(&a)
    
    f.Println("Digite um número inteiro \"B\": ")
    f.Scan(&b)
    f.Println("\n")
    
    if a % b == 0 {
        f.Println("O número A = ", a, " é divisível pelo número B = ", b )
    } else {
        f.Println("Os números não são divisíveis um pelo outro.")
    }
}
