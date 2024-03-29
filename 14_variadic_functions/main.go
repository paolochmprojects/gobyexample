package main

import "fmt"

func sumInts(nums ...int) int {
	var total int
	for _, value := range nums {
		total += value
	}
	return total
}

func main() {
	res := sumInts(1, 2, 3, 4, 5, 6)
	fmt.Println("sum total =", res)

	sliceNum := []int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10, 11, 12, 13}

	res2 := sumInts(sliceNum...)
	fmt.Println("sum total 2 =", res2)
}
