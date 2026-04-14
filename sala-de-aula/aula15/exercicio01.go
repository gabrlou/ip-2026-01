package main
import "fmt"
func main() {
	fmt.Println(potencia(2,4))
}

func potencia (x, n int) int {
	if n == 0 {
		return 1
	}
	return x * potencia(x, n-1)
}
