package main
import "fmt"

func main() {
    
    var nota, soma, alunos, i, media float64
    
    fmt.Println("Qual a quantidade de alunos da turma? ")
    fmt.Scan(&alunos)
    fmt.Println("\n")
    
    soma = 0
    i = 1
  for  i <= alunos {
     fmt.Println("Qual a nota do ", i, "º aluno? ")
     fmt.Scan(&nota)
     
     i++
     soma += nota
     }
  
  media = soma / alunos
  fmt.Println("\nA média da turma é: ", fmt.Sprintf("%.2f",media))
}
