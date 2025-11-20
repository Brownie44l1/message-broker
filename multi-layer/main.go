package main

import (
	"errors"
	"fmt"
)

type RoutingLayer struct {
	Topics       map[string][]string
	Subscriber   map[string]map[string]int
	Subscription map[string]map[string]bool
}

func (r *RoutingLayer) AddTopic(topic string) error {
	if _, exists := r.Topics[topic]; exists {
		return errors.New("topic exists")
	}
	r.Topics[topic] = []string{}
	return nil
}

func (r *RoutingLayer) AddSubscriber(name string) error {
	if _, exists := r.Subscriber[name]; exists {
		return errors.New("subscriber exists")
	}
	r.Subscriber[name] = map[string]int{}
	r.Subscription[name] = map[string]bool{}
	return nil
}

func (r *RoutingLayer) Subscribe(subscriber, topic string) error {
	if _, exists := r.Subscriber[subscriber]; !exists {
		return errors.New("subscriber does not exists")
	}
	if _, exists := r.Topics[topic]; !exists {
		return errors.New("topic does not exists")
	}
	r.Subscription[subscriber][topic] = true
	r.Subscriber[subscriber][topic] = 0
	return nil
}

func (r *RoutingLayer) Publish(topics []string, msg string) error {
	for _, topic := range topics {
		if _, exists := r.Topics[topic]; !exists {
			return errors.New("topic does not exists")
		}
		r.Topics[topic] = append(r.Topics[topic], msg)
	}
	return nil
}

func (r *RoutingLayer) Consume(subscriber string) ([]string, error) {
	if _, exists := r.Subscriber[subscriber]; !exists {
		return nil, errors.New("subscriber does not exists")
	}

	var messages []string

	for topic := range r.Subscription[subscriber] {
		if r.Subscription[subscriber][topic] {
			offset := r.Subscriber[subscriber][topic]
			msg := r.Topics[topic]
			if offset < len(msg) {
				messages = append(messages, msg[offset])
				r.Subscriber[subscriber][topic] = offset + 1
			}
		}
	}
	return messages, nil
}

func main() {
	router := &RoutingLayer{
		Topics:       make(map[string][]string),
		Subscriber:   make(map[string]map[string]int),
		Subscription: make(map[string]map[string]bool),
	}

	_ = router.AddTopic("orders")
	_ = router.AddTopic("logs")

	_ = router.AddSubscriber("api-1")
	_ = router.AddSubscriber("api-2")

	_ = router.Subscribe("api-1", "orders")
	_ = router.Subscribe("api-2", "logs")

	topics := []string{"orders", "logs"}
	_ = router.Publish(topics, "testing")

	messages, _ := router.Consume("api-1")
	for _, msg := range messages {
		fmt.Println(msg)
	}
}
