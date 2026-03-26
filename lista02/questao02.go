package main
import  "fmt"

func main() {
    
    var numero int
    
    fmt.Println("Digite um número inteiro: ")
    fmt.Scan(&numero)
    fmt.Println("\n")

    if numero < 0 {
        fmt.Println("O número digitado é negativo.")
    } else if numero > 0 {
         fmt.Println("O número digitado é positivo.")
    } else {
        fmt.Println("O número digitado é nulo.")
    }

}
