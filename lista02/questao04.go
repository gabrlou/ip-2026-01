package main
import (
    "fmt"
    "math"
)

func main() {
    
    var numero, raiz, quadrado float64
    
    fmt.Println("Digite um número do tipo float: ")
    fmt.Scan(&numero)
    fmt.Println("\n")
    
    if numero > 0 || numero == 0 {
        raiz = math.Sqrt(numero)
        fmt.Printf("A raiz quadrada do número é: %.2f\n", raiz)
        
    } else {
        quadrado = numero * numero
        fmt.Println("O quadrado do número é: ", quadrado)
    }

}
