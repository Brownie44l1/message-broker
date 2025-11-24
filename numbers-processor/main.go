package main

import "fmt"

type Number struct {
	Original int
	Squared  int
}

func generator(out chan<- Number) {
	for i := 1; i <= 10; i++ {
		out <- Number{Original: i, Squared: 0}
	}
	close(out)
}

func processor(in <-chan Number, out chan<- Number) {
	for n := range in {
		n.Squared = n.Original * n.Original
		out <- n
	}
	close(out)
}

func printer(in <-chan Number, out chan<- Number) {
	for n := range in {
		fmt.Printf("Received: %d, Squared: %d\n", n.Original, n.Squared)
		out <- n
	}
	close(out)
}

func sum(in <-chan Number) {
	total := 0
	for n := range in {
		total += n.Squared
	}
	fmt.Printf("Total sum of the squared numbers is: %d\n", total)
}

func main() {
	ch1 := make(chan Number)
	ch2 := make(chan Number)
	ch3 := make(chan Number)

	go generator(ch1)
	go processor(ch1, ch2)
	go printer(ch2, ch3)
	sum(ch3)

	fmt.Println("Refactored: all done.")
}
