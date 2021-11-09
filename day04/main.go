package main

import (
	"fmt"
)

func main() {
	mapA := make(map[string]int)
	mapA["k1"] = 12
	mapA["k2"] = 13
	fmt.Println("mapA:", mapA)
	mapB := map[string]int{
		"k1": 12,
		"k2": 13,
	}

	fmt.Println("mapB:", mapB)
	fmt.Println("Deleting key `k1` from mapB...")
	delete(mapB, "k1")
	delete(mapB, "k1")
	fmt.Println("mapB:", mapB)
	fmt.Println()
}