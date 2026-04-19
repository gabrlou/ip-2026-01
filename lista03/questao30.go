package main

import (
	f "fmt"
	m "math"
)

func main() {
	var volume float64

	for r := 0.0; r <= 20.0; r += 0.5 {

		volume = (4.0 / 3.0) * m.Pi * m.Pow(r, 3)

		f.Printf("Raio: %.1f cm | Volume: %.2f cm³\n", r, volume)
	}
}
