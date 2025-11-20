package main

import (
	"errors"
	"fmt"
)

type MessageBroker struct {
	Topics map[string][]string
}

func (m *MessageBroker) CreateTopic(topic string) {
	m.Topics[topic] = []string{}
}

func (m *MessageBroker) Publish(topic string, message string) error {
	if _, exist := m.Topics[topic]; !exist {
		return errors.New("The topic does not exist")
	}
	m.Topics[topic] = append(m.Topics[topic], message)
	return nil
}

func (m *MessageBroker) Consume(topic string) (string, error) {
	messages, exists := m.Topics[topic]
	if !exists {
		return "", errors.New("topic does not exist")
	}
	
	if len(m.Topics[topic]) == 0 {
		return "", errors.New("topic is empty")
	}

	msg := messages[0]
	m.Topics[topic] = messages[1:]
	return msg, nil
}

func main() {
	broker := &MessageBroker{
		Topics: make(map[string][]string),
	}

	broker.CreateTopic("Orders")
	_ = broker.Publish("Orders", "2 kay 4 d way")
	_ = broker.Publish("Orders", "eggroll for lunch")

	message, err := broker.Consume("Orders")
	if err != nil {
		fmt.Println("Error: ", err)
	} else {
		fmt.Println("Messages: ", message)
	}
}