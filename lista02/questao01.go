package main
import  "fmt"

func main() {
    
    var numero int
    
    fmt.Println("Digite um número inteiro: ")
    fmt.Scan(&numero)
    fmt.Println("\n")

    if numero % 2 == 0 {
        fmt.Println("O número digitado é par.")
    } else {
         fmt.Println("O número digitado é ímpar.")
    }

}
