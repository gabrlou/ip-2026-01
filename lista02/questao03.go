package main
import  "fmt"

func main() {
    
    var n1, n2, soma int
    
    fmt.Println("Digite um primeiro número inteiro: ")
    fmt.Scan(&n1)
    
    fmt.Println("Digite um segundo número inteiro: ")
    fmt.Scan(&n2)
    fmt.Println("\n")
    
    soma = n1 + n2
    if soma > 20 {
        fmt.Println("O resultado é: ", soma + 8)
    } else {
         fmt.Println("O resultado é: ", soma - 5)
    }

}
