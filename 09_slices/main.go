package main

import (
	"fmt"
	"slices"
)

func main() {
	var a []string

	fmt.Println("uninit", a, a == nil, len(a) == 0)

	s := make([]string, 4)
	fmt.Println("emp:", s, "len:", len(s), "cap:", cap(s))

	s[0] = "a"
	s[1] = "b"
	s[2] = "c"
	s[3] = "d"

	fmt.Println("set:", s)
	fmt.Println("get:", s[3])
	fmt.Println("len:", len(s))

	s = append(s, "f")
	s = append(s, "g", "h")
	fmt.Println("apn:", s)

	c := make([]string, len(s))
	copy(c, s)
	fmt.Println("copy:", c)

	l := s[2:5]
	fmt.Println("sl1:", l)

	l = s[:5]
	fmt.Println("sl2:", l)

	l = s[2:]
	fmt.Println("sl3:", l)

	t := []string{"g", "h", "i"}
	fmt.Println("dcl:", t)

	t2 := []string{"g", "h", "i"}
	fmt.Println("dcl:", t2)

	if slices.Equal(t, t2) {
		fmt.Println("t and t2 are equal.")
	}

	twoD := make([][]int, 3)
	for i := 0; i < 3; i++ {
		innerLen := i + 1
		twoD[i] = make([]int, innerLen)
		for j := 0; j < innerLen; j++ {
			twoD[i][j] = i + j
		}
	}
	fmt.Println("2d: ", twoD)

	var arr3 []int

	for i := 0; i < 100; i++ {
		arr3 = append(arr3, i)
	}
	fmt.Println(arr3)
}
