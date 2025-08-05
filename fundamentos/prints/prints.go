package main

import "fmt"

func main() {
	fmt.Println("Mesma")
	fmt.Println("linha")

	fmt.Println("Nova linha")
	fmt.Print("Nova ")

	x := 3.141516

	xs := fmt.Sprintf("%.6f", x)

	fmt.Println("O valor de x é " + xs)
	fmt.Printf("O valor de x é %.2f\n", x) // Formatação com printf
}
