package main

import "fmt"

func main() {
	// declaracion de una variable.
	var a = "string"
	fmt.Println(a)

	// declaracion multiple.
	var b, c = 1, 2
	fmt.Println(b, c)

	// inferencia de una variable
	var d = true
	fmt.Println(d)

	// por defecto se le asigna valor 0 a un tipo de variable int
	var e int
	fmt.Println(e)
	// por defecto se le asigna valor "" a un tipo de variable string
	var f string
	fmt.Println(f)
	// por defecto se le asigna valor false a un tipo de variable bool
	var g bool
	fmt.Println(g)

	// Es una abreviatura para declarar e inicializar una variable
	h := "new string"
	fmt.Println(h)
}
