package main

import (
	"errors"
	"fmt"
)

type MessageBroker struct {
	Topics        map[string]map[string]int
	TopicsMessage map[string][]string
}

func (m *MessageBroker) CreateTopic(topic string) error {
	if _, exists := m.Topics[topic]; exists {
		return errors.New("topic exists")
	}
	m.Topics[topic] = map[string]int{}
	m.TopicsMessage[topic] = []string{}
	return nil
}

func (m *MessageBroker) Subscribe(subscriberID string, topic string) error {
	if _, exists := m.Topics[topic]; !exists {
		return errors.New("topic does not exist")
	}

	if _, subscriberExists := m.Topics[topic][subscriberID]; subscriberExists {
		return errors.New("subscriber already exist for this topic")
	}
	m.Topics[topic][subscriberID] = 0
	return nil
}

func (m *MessageBroker) Publish(topic string, message string) error {
	m.TopicsMessage[topic] = append(m.TopicsMessage[topic], message)
	return nil
}

func (m *MessageBroker) Consume(subscriberID string, topic string) (string, error) {
	if _, exists := m.Topics[topic]; !exists {
		return "", errors.New("topic does not exist")
	}

	readIndex, exists := m.Topics[topic][subscriberID]
	if !exists {
		return "", errors.New("subscriber does not exist")
	}

	message := m.TopicsMessage[topic]
	if len(message) == 0 {
		return "", errors.New("no message available for this topic")
	}
	if readIndex >= len(message) {
		return "", errors.New("no new message")
	}

	msg := message[readIndex]
	m.Topics[topic][subscriberID] = readIndex + 1
	return msg, nil
}

func main() {
	broker := &MessageBroker{
		Topics:        make(map[string]map[string]int),
		TopicsMessage: make(map[string][]string),
	}

	_ = broker.CreateTopic("logs")
	_ = broker.CreateTopic("orders")

	_ = broker.Subscribe("error-log", "logs")
	_ = broker.Subscribe("pending-order", "orders")
	_ = broker.Subscribe("service-A", "orders")
	_ = broker.Subscribe("service-B", "orders")

	_ = broker.Publish("logs", "10mins downtime, 40mins uptime")
	_ = broker.Publish("orders", "pizza delivery to Lekki")
	_ = broker.Publish("orders", "Order #1")
	_ = broker.Publish("orders", "Order #2")

	msg1, _ := broker.Consume("service-A", "orders")
	fmt.Println("Service A:", msg1)

	msg2, _ := broker.Consume("service-B", "orders")
	fmt.Println("Service B:", msg2)

	msg3, _ := broker.Consume("service-A", "orders")
	fmt.Println("Service A:", msg3)

	msg4, _ := broker.Consume("service-B", "orders")
	fmt.Println("Service B:", msg4)

	msg5, err := broker.Consume("service-A", "orders")
	if err != nil {
		fmt.Println("Service A caught up:", err)
	} else {
		fmt.Println("Service A:", msg5)
	}

	msg6, _ := broker.Consume("error-log", "logs")
	fmt.Println(msg6)
	msg7, _ := broker.Consume("pending-order", "orders")
	fmt.Println(msg7)
}
