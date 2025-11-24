package main

import (
	"fmt"
	"sync"
	"time"
)

func worker(id int, duration time.Duration) {
	fmt.Printf("Worker %d starting, (will work for %s)\n", id, duration)
	time.Sleep(duration)
	fmt.Printf("Worker %d done.\n", id)
}

func main() {
	var wg sync.WaitGroup
	start := time.Now()

	wg.Add(5)
	go func () {
		worker(1, 2*time.Millisecond)
		wg.Done()
	}()
	go func () {
		worker(2, 3*time.Millisecond)
		wg.Done()
	}()
	go func () {
		worker(3, 1*time.Millisecond)
		wg.Done()
	}()
	go func () {
		worker(4, 3*time.Millisecond)
		wg.Done()
	}()
	go func () {
		worker(5, 1*time.Millisecond)
		wg.Done()
	}()

	wg.Wait()
	elapsed := time.Since(start)
	fmt.Println(elapsed)
}