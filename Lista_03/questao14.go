package main

import (
	"fmt"
	f "fmt"
)

func main() {
	var n1, n2 int

	f.Print("\nDigite o valor de N1: ")
	f.Scan(&n1)

	f.Print("Digite o valor de N2: ")
	f.Scan(&n2)

	if n1 > n2 {
		n1, n2 = n2, n1
	}

	f.Printf("\nNúmeros primos entre %d e %d:\n", n1, n2)

	for i := n1; i <= n2; i++ {
		if i < 2 {
			continue
		}

		primo := true
		for j := 2; j*j <= i; j++ {
			if i%j == 0 {
				primo = false
				break
			}
		}

		if primo == true {
			fmt.Printf("%d ", i)
		}
	}
	f.Printf("\n\n")
}
