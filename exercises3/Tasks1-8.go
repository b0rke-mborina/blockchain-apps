package main

import (
	"fmt"
	"math/rand"
	"sort"
	"strconv"
	"strings"
)

func main() {
	// 1.
	// arr := []string{"4", "5", "3"}
	arr := []string{"2", "3", "-1", "34", "53"}
	var arr2 []int
	for _, i := range arr {
		j, err := strconv.Atoi(i)
		if err != nil {
			panic(err)
		}
		arr2 = append(arr2, j)
	}
	sort.Ints(arr2)
	fmt.Println("Array:", arr2)
	fmt.Println()

	// 2.
	var nums []int
	for i := 0; i < 10; i++ {
		// x, _ := fmt.Scanf("%d", &i)
		var x int
		fmt.Scan(&x)
		nums = append(nums, x)
	}
	fmt.Println(nums)
	fmt.Println(sum(nums))
	fmt.Println()

	// 3.
	l := rand.Intn(10)
	var rands []int
	for i := 0; i < l; i++ {
		x := rand.Intn(10)
		rands = append(rands, x)
	}
	rands = removeDuplicateInt(rands)
	sort.Ints(rands)
	fmt.Println(rands)
	fmt.Println()

	// 4.
	reqs := []int{
		500, 600, 250, // 1. dan
		900, 800, 600, // 2. dan
		150, 654, 235, // 3. dan
		121, 876, 285, // 4. dan
	}
	chunkSize := 3
	for i := 0; i < len(reqs); i += chunkSize {
		x := reqs[i] + reqs[i+1] + reqs[i+2]
		fmt.Println("Dan", i / chunkSize + 1, "=", x)
	}
	fmt.Println()

	// 5.
	rooms := map[int][]string {
		101: {"Ivica", "Marko", "Teo"},
		103: {"Mirna", "Barbara", "Petra"},
		403: {"Antonio", "Karlo", "Karlo"},
		406: {"Marko", "Ivo", "Bobo"},
		1:   {"recepcija"},
	}
	delete(rooms, 1)
	for k, v := range rooms {
		fmt.Println("Room", k, "=", strings.Join(v[:], ", "))
	}
	fmt.Println()

	// 6.
	m1 := map[int][]int {
		1: {3, 22, 1},
		2: {7, 32, 420},
		52: {423, 42, 43},
	}
	fmt.Println("m1 poklapanja:")
	printMatches(m1)
	m2 := map[int][]int {
		2: {42, 32, 7},
		1: {22, 3, 1},
		520: {425, 42, 43},
	}
	fmt.Println("m2 poklapanja:")
	printMatches(m2)

	// 7.
	text := "This is some text with many words and text which is a string"
	words := strings.Split(text ," ")
	wordCounts := make(map[string]int)
	for _, word := range words {
		if count, ok := wordCounts[word]; ok {
			wordCounts[word] = count + 1
		} else {
			wordCounts[word] = 1
		}
	}
	fmt.Println("Word counts:", wordCounts)
	fmt.Println()

	// 8.
	temps := map[string][]int {
		"Pula": {10, 6, 11, 4, 6, 5, 7},
		"Zagreb": {13, 12, 12, 9, 11, 8},
		"Rijeka": {8, 5, 7, 3, 5, 4, 6},
	}
	tempsPula := temps["Pula"]
	fmt.Println("Average temperature in Pula:", sum(tempsPula) / len(tempsPula))
	fmt.Println()
}

func sum(numbs []int) int {
	result := 0
	for i := 0; i < len(numbs); i++ {
		result += numbs[i]
	}
	return result
}

func removeDuplicateInt(intSlice []int) []int {
	allKeys := make(map[int]bool)
	var list []int
	for _, item := range intSlice {
		if _, value := allKeys[item]; !value {
			allKeys[item] = true
			list = append(list, item)
		}
	}
	return list
}

func printMatches(m map[int][]int) {
	for k, v := range m {
		for _, el := range v {
			if el == k {
				fmt.Println("k:", k, "el:", el)
			}
		}
	}
	fmt.Println()
}
