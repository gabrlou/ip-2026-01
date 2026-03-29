package main

import f "fmt"

func main() {
	var x, fx float64

	f.Print("Digite um valor para \"X\": ")
	f.Scan(&x)

	fx = 8 / (2 - x)

	f.Printf("\nO valor de \"f(x) = 8/(2-x)\" é: %.2f\n", fx)
}
