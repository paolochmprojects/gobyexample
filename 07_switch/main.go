package main

import (
	"fmt"
	"time"
)

func main() {
	i := 2

	fmt.Print("Escribe ", i, " como ")
	switch i {
	case 1:
		fmt.Println("One")
	case 2:
		fmt.Println("two")
	case 3:
		fmt.Println("three")
	}

	switch time.Now().Weekday() {
	case time.Saturday, time.Sunday:
		fmt.Println("It's a weekend.")
	default:
		fmt.Println("It's a weekday.")
	}

	t := time.Now()
	switch {
	case t.Hour() < 12:
		fmt.Println("Is before noon")
	default:
		fmt.Println("Is after noon")
	}

	whatAmI := func(i interface{}) {
		switch t := i.(type) {
		case bool:
			fmt.Println("I'am bool.")
		case int:
			fmt.Println("I'am int.")
		default:
			fmt.Printf("Don't know type %T\n", t)
		}
	}

	whatAmI(true)
	whatAmI(23)
	whatAmI("Hola")
}
