package main

import f "fmt"

func main() {
	const N = 50
	fib := make([]int, N)

	fib[0] = 1
	fib[1] = 1

	for i := 2; i < N; i++ {
		fib[i] = fib[i-1] + fib[i-2]
	}

	f.Print("Série de Fibonacci:\n")
	for i := 0; i < N; i++ {
		if i == 0 {
			f.Printf("%d", fib[i])
		} else {
			f.Printf(" - %d", fib[i])
		}
	}

	f.Print("\n\n")
}
