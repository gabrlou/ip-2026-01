package main

import f "fmt"

func main() {
	var (
		d1, d2 int
		soma   float64
	)

	d1 = 37
	d2 = 38
	soma = 0
	for i := 1; i <= 37; i++ {
		soma += float64(d1*d2) / float64(i)
		d1--
		d2--
	}
	f.Printf("\nResultado da soma [S = ((37 * 38) / 1) + ((36 * 37) / 2) ...]: %.2f\n\n", soma)
}
