package main
import "fmt"
    func main() {

var l1, l2, l3 float64

fmt.Println("Digite o lado 1 do triângulo: ")
fmt.Scan(&l1)

fmt.Println("Digite o lado 2 do triângulo: ")
fmt.Scan(&l2)

fmt.Println("Digite o lado 3 do triângulo: ")
fmt.Scan(&l3)

if (l1 + l2) > l3 || (l1 + l3) > l2 || (l2 +l3) > l1 {

    if l1 == l2 && l1 == l3 {
        fmt.Println("\nO triângulo é equilátero!")
    } else if l1 == l2 || l1 == l3 || l2 == l3 {
        fmt.Println("\nO triângulo é isósceles!")
    } else {
        fmt.Println("\nO triângulo é escaleno!")
    }

} else {
    fmt.Println("Triângulo Inválido!")
}
}
