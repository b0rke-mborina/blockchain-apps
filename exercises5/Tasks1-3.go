package main

import (
	"fmt"
	"math/rand"
	"sync"
	"time"
)

// 1.
func routine(chnl chan int) {
	l := rand.Intn(100 - 1) + 1 // (max - min) + min
	for i := 0; i < l; i++ {
		x := rand.Intn(100 - 1) + 1
		if x % 2 == 0 {
			chnl <- x
		}
	}
	close(chnl)
}

// 2.
func goroutina(routineNumber int, wg *sync.WaitGroup) {
	numberOfSeconds := rand.Intn(10 - 1) + 1
	time.Sleep(time.Duration(numberOfSeconds) * time.Second)
	fmt.Println("Goroutina", routineNumber, "je završila.")
	wg.Done()
}

// 3.
func writeNumbers(chnl chan int) {
	for i := 1; i <= 100; i++ {
		chnl <- i
	}
	close(chnl)
}
func readNumbers(chnl chan int) {
	for v := range chnl {
		fmt.Println("Ispis iz kanala 2:", v)
	}
}

func main() {
	// 1.
	ch := make(chan int)
	go routine(ch)
	for v := range ch {
		fmt.Println("Spremljeno:", v)
	}
	fmt.Println()

	// 2.
	var wg sync.WaitGroup
	for i := 0; i < 100; i++ {
		wg.Add(1)
		go goroutina(i + 1, &wg)
	}
	wg.Wait()
	fmt.Println("Kraj programa")

	// 3.
	chan1 := make(chan int, 10)
	chan2 := make(chan int, 20)
	go writeNumbers(chan1)
	go readNumbers(chan2)
	/*for v := range chan1 {
		if v == 100 {
			close(chan1)
		}
		if rand.Intn(2) == 1 {
			fmt.Println("Ispis iz kanala 1:", v)
		} else {
			chan2 <- rand.Intn(10000 - 1000) + 1000
		}
	}*/
	counter := 0
	petlja:
	for {
		select {
			// čekaj na vrijednost u output1
			case s1, ok := <-chan1:
				if ok {
					fmt.Println(s1)
				}
			// čekaj na vrijednost u output2
			case chan2 <- rand.Intn(10000 - 1000) + 1000:
				fmt.Println("Printing...")
				if counter == 100 {
					break petlja
				}
				counter++
			default:
				break
		}
	}
}
