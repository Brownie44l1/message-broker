package main

import (
	"fmt"
	"time"
)

func main() {
    ch1 := make(chan int)
    ch2 := make(chan int)
    
    go func() {
        ch1 <- 1
        val := <-ch2
        fmt.Println("Goroutine got:", val)
    }()
    
    val := <-ch1
    fmt.Println("Main got:", val)
	time.Sleep(2 * time.Millisecond)
	ch2 <- 2
}