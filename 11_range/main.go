package main

import "fmt"

func main() {
	numSlice := []int{}
	numEven := []int{}
	var sum int

	for i := 0; i <= 200; i++ {
		numSlice = append(numSlice, i)
	}

	for _, val := range numSlice {
		if val%2 == 0 {
			numEven = append(numEven, val)
		}
		sum += val
	}

	fmt.Println("sum:", sum)
	fmt.Println("Even numbers:", numEven)

	kvs := map[string]string{
		"fruit1": "banana",
		"fruit2": "apple",
		"fruit3": "blueberry",
	}

	for key, value := range kvs {
		fmt.Printf("key: %s -> value: %s\n", key, value)
	}

	for key := range kvs {
		println("key:", key)
	}

	for index, character := range "Paolo" {
		fmt.Println(index, character)
	}
}
