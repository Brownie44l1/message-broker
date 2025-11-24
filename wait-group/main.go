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

	durations := []time.Duration{
		2 * time.Millisecond,
		3 * time.Millisecond,
		1 * time.Millisecond,
		3 * time.Millisecond,
		1 * time.Millisecond,
	}

	for i, duration := range durations {
		wg.Add(1)

		go func(workerID int, sleepDuration time.Duration)  {
			defer wg.Done()
			worker(workerID, sleepDuration)
		}(i+1, duration)
	}

	wg.Wait()
	elapsed := time.Since(start)
	fmt.Println(elapsed)
}