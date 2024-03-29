package main

import "fmt"

func main() {

	var i int = 1

	fmt.Println("Primer for")
	for i <= 4 {
		fmt.Println(i)
		// operacion equivalente i = i + 1
		i += 1
	}

	fmt.Println("Segundo for")
	for j := 0; j <= 10; j++ {
		fmt.Println(j)
	}

	fmt.Println("Tercer for")
	for {
		println("Primera iteracion")
		// ruptura de bucle infinito
		break
	}

	for k := 0; k <= 10; k++ {
		if k%2 == 0 {
			fmt.Println("Omision de iteracion")
			continue
		}
		fmt.Println(k)
	}
}
