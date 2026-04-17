package main

import f "fmt"

func main() {
	var n1, n2 int

	f.Print("\nDigite o valor de N1: ")
	f.Scan(&n1)

	f.Print("Digite o valor de N2: ")
	f.Scan(&n2)

	resultado := 0
	original_n1 := n1
	original_n2 := n2

	if n2 < 0 {
		n2 = -n2
		n1 = -n1
	}

	for i := 0; i < n2; i++ {
		resultado += n1
	}

	if original_n1 < 0 && original_n2 < 0 {
		f.Printf("\n	(%d) * (%d) = %d\n\n", original_n1, original_n2, resultado)
	} else if original_n1 < 0 && original_n2 >= 0 {
		f.Printf("\n	(%d) * %d = %d\n\n", original_n1, original_n2, resultado)
	} else if original_n2 < 0 && original_n1 >= 0 {
		f.Printf("\n	%d * (%d) = %d\n\n", original_n1, original_n2, resultado)
	} else {
		f.Printf("\n	%d * %d = %d\n\n", original_n1, original_n2, resultado)
	}
}
