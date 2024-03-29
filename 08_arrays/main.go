package main

import "fmt"

func main() {
	var arr1 [5]int
	fmt.Println("em:", arr1) //em: [0 0 0 0 0]

	arr1[4] = 100
	fmt.Println("set:", arr1)
	fmt.Println("get:", arr1[4])

	fmt.Println("length of arr1 =", len(arr1))

	b := [5]int{1, 2, 3, 4, 5}
	fmt.Println("dcl:", b)

	// array multidimensionales
	var twoD [2][3]int
	for i := 0; i < 2; i++ {
		for j := 0; j < 3; j++ {
			twoD[i][j] = i + j
		}
	}
	fmt.Println("2d: ", twoD)
}
