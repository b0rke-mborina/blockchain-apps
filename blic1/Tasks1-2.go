package main

import (
	"fmt"
	"math/rand"
)

func main() {
	niz := make([]int, 0)
	for i := 0; i < 10; i++ {
		x := rand.Intn(100)
		niz = append(niz, x)
	}
	fmt.Println(max(niz))
}

func max(arr []int) (int, int) {
	result := arr[0]
	index := 0
	for i := 1; i < len(arr); i++ {
		if arr[i] > result {
			result = arr[i]
			index = i
		}
	}
	return result, index
}
