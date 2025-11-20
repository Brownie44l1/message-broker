package main

import (
	"errors"
	"fmt"
)

type MessageBroker struct {
	messages []string
}

func (m *MessageBroker) publish(message string) {
	m.messages = append(m.messages, message)
}

func (m *MessageBroker) consume() (string, error) {
	if len(m.messages) == 0 {
		return "", errors.New("Empty list")
	} else {
		msg := m.messages[0]
		m.messages = m.messages[1:]
		return msg, nil
	}
}

func main() {
	broker := &MessageBroker{}

	broker.publish("Hello")
	broker.publish("World")

	msg, err := broker.consume()

	if err != nil {
		fmt.Println("Error: ", err)
	} else {
		fmt.Println("Got Message: ", msg)
	}
}
