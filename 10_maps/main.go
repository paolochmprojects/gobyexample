package main

import (
	"fmt"
	"maps"
)

func main() {
	map1 := make(map[string]int)

	map1["k1"] = 1
	map1["k2"] = 2

	fmt.Println("map:", map1)

	v1 := map1["k1"]
	fmt.Println("v1:", v1)

	v3 := map1["k3"]
	fmt.Println("v3", v3)

	fmt.Println(len(map1))

	delete(map1, "k2")
	fmt.Println("map:", map1)

	clear(map1)
	fmt.Println("map", map1)

	dd, prs := map1["k1"]
	fmt.Println(dd, prs)

	map2 := map[string]int{"foo": 1, "bar": 2}

	val, prss := map2["bar"]
	fmt.Println(val, prss)

	map3 := map[string]int{"foo": 1, "bar": 2}
	if maps.Equal(map2, map3) {
		fmt.Println("map1 and map2 are equal")
	}
}
