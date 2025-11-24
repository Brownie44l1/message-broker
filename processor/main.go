package main

import (
	"fmt"
	"time"
)

type Transaction struct {
	ID     int
	Amount int
	Type   string
}

func ProcessTeller(tellerName string, ch <-chan Transaction, resultCH chan<- string) {
	for transaction := range ch {
		time.Sleep(500 * time.Millisecond)
		message := fmt.Sprintf("%s processed %s of $%d (ID: %d)",
			tellerName, transaction.Type, transaction.Amount, transaction.ID)

		resultCH <- message
	}
}

func main() {
	ch := make(chan Transaction)
	resultCH := make(chan string)

	transactions := []Transaction{
		{
			ID:     1,
			Amount: 100,
			Type:   "deposit",
		},
		{
			ID:     2,
			Amount: 50,
			Type:   "withdrawal",
		},
		{
			ID:     3,
			Amount: 200,
			Type:   "deposit",
		},
		{
			ID:     4,
			Amount: 75,
			Type:   "withdrawal",
		},
		{
			ID:     5,
			Amount: 150,
			Type:   "deposit",
		},
		{
			ID:     6,
			Amount: 30,
			Type:   "withdrawal",
		},
	}

	go ProcessTeller("Teller-A", ch, resultCH)
	go ProcessTeller("Teller-B", ch, resultCH)

	go func() {
		for _, transaction := range transactions {
			ch <- transaction
		}
		close(ch)
	}()

	for i := 0; i < 6; i++ {
		msg := <-resultCH
		fmt.Println(msg)
	}

	fmt.Println("All transactions processed!")
}
