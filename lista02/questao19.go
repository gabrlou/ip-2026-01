package main

import (
	f "fmt"
	m "math"
)

func main() {
	var (
		figura                     int
		volume, area, altura, raio float64
	)

	f.Print("Digite o tipo de figura a ser calculada (1 - cone / 2 - cilindro / 3 - esfera): ")
	f.Scanln(&figura)
	if figura < 1 || figura > 3 {
		f.Println("Valor inválido!\nDigite (1 - cone / 2 - cilindro / 3 - esfera).")
		return
	}

	if figura == 1 { //Cone
		f.Print("\nDigite a altura do cone (em metros): ")
		f.Scanln(&altura)

		f.Print("\nDigite o raio do cone (em metros): ")
		f.Scanln(&raio)

		volume = (m.Pi * m.Pow(raio, 2) * altura) / 3
		area = m.Pi * raio * (raio + m.Sqrt(m.Pow(raio, 2)+m.Pow(altura, 2)))

		f.Print("\n")
		f.Printf("Figura: Cone Reto\nVolume: %.2f metros cúbicos.\nÁrea: %.2f metros quadrados.", volume, area)
	} else if figura == 2 { //Cilindro
		f.Print("\nDigite a altura do cilindro (em metros): ")
		f.Scanln(&altura)

		f.Print("\nDigite o raio do cilindro (em metros): ")
		f.Scanln(&raio)

		volume = m.Pi * m.Pow(raio, 2) * altura
		area = 2 * m.Pi * raio * (raio + altura)

		f.Print("\n")
		f.Printf("Figura: Cilindro\nVolume: %.2f metros cúbicos.\nÁrea: %.2f metros quadrados.", volume, area)
	} else { //Esfera
		f.Print("\nDigite o raio da esfera (em metros): ")
		f.Scanln(&raio)

		volume = (4 * m.Pi * m.Pow(raio, 3)) / 3
		area = 4 * m.Pi * m.Pow(raio, 2)

		f.Print("\n")
		f.Printf("Figura: Esfera\nVolume: %.2f metros cúbicos.\nÁrea: %.2f metros quadrados.", volume, area)
	}
}
