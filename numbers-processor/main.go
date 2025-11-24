package main

import (
	"fmt"
)

func generator(out chan<- int) {
	for i := 1; i <= 10; i++ {
		out <- i
	}
	close(out)
}

func receiver(in <- chan int, out chan <- int) {
	for v := range in {
		out <- v * v
	}
	close(out)
}

func printer(in <- chan int, out chan <- int) {
	for squared := range in {
		original := 0
		for i := 1; i*i <= squared; i++ {
			if i*i == squared {
				original = i
				break
			}
		}
		fmt.Printf("Received: %d -> Squared: %d\n", original, squared)
		out <- squared
	}
	close(out)
}

func sum(in <- chan int) {
	var sum int
	for v := range in {
		sum += v
	}
	fmt.Printf("The total sum of squared results is: %d\n", sum)
}

func main() {
	a := make(chan int)
	b := make(chan int)
	c := make(chan int)

	go generator(a)
	go receiver(a, b)
	go printer(b, c)
	sum(c)

	fmt.Println("All done")
}