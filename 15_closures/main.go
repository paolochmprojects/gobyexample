package main

import "fmt"

func intSeq() func() int {
	i := 0
	return func() int {
		i++
		return i
	}
}

func main() {

	nextStep := intSeq()

	fmt.Println(nextStep())
	fmt.Println(nextStep())
	fmt.Println(nextStep())
	fmt.Println(nextStep())

	newNextStep := intSeq()
	fmt.Println(newNextStep())
}
