package main

import (
	"bufio"
	"fmt"
	"math/rand"
	"os"
	"strings"
)

func main() {
	// 1.
	result := 0
	for i := 1; i <= 1000; i++ {
		if i % 2 == 0 {
			result += i
			fmt.Print(i, " ")
			if i != 1000 {
				print("+ ")
			}
		}
	}
	fmt.Print("= ", result, "\n\n")

	// 2.
	// charset := "\\|/-"
	// for {
		// fmt.Println(string(charset[rand.Intn(len(charset))]), " Please wait. Processing...")
	// }

	// 3.
	value := rand.Intn(10)
	fmt.Println("User value: ", value)
	y := 0
	petlja:
		for i := 0; i < 5; i++ {
			x := rand.Intn(10)
			fmt.Println(x)
			if value == x && i == 0 {
				fmt.Println("Pobjednik isprve! WOW!!!\n")
				y = 1
				break petlja
			} else if value == x {
				fmt.Println("Bravo, pogodio si broj\n")
				y = 1
				break petlja
			}
		}
	if y == 0 {
		fmt.Println("Pokušaj ponovo :(\n")
	}

	// 4.
	str, strLen := f()
	fmt.Println(str)
	fmt.Println(strLen)
}

func f() (string, int) {
	fmt.Print("Enter something: ")
	reader := bufio.NewReader(os.Stdin)
	input, err := reader.ReadString('\n')
	if err != nil {
		fmt.Println("An error occured while reading input. Please try again", err)
	}
	inputVal := strings.TrimSpace(input)
	return inputVal, len(inputVal)
}
