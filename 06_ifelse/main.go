package main

import "fmt"

func main() {

	if 7%2 == 0 {
		fmt.Println("7 es par")
	} else {
		fmt.Println("7 es impar")
	}

	if 8%4 == 0 {
		fmt.Println("8 es divisible por 4")
	}

	if 8%2 == 0 || 7%2 == 0 {
		fmt.Println("8 o 7 es par")
	}

	if num := 0; num < 0 {
		fmt.Println("numero es negativo")
	} else if num < 10 {
		fmt.Println("numero tiene uno solo digito")
	} else {
		fmt.Println("numero tiene multiples digitos")
	}
}
