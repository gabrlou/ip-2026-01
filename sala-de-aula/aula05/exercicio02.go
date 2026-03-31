package main
import "fmt"
    func main() {

var n1, n2, p1, p2, media_pond float64

fmt.Println("Digite a primeira nota: ")
fmt.Scan(&n1)

fmt.Println("Digite o peso da primeira nota: ")
fmt.Scan(&p1)

fmt.Println("Digite a segunda nota: ")
fmt.Scan(&n2)

fmt.Println("Digite o peso da segunda nota: ")
fmt.Scan(&p2)

media_pond = ((n1 * p1) + (n2 * p2)) / (p1 + p2)

fmt.Println("\nA média ponderada é: ", fmt.Sprintf("%.2f", media_pond))

}
