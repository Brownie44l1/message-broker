package main

import (
	"errors"
	"fmt"
)

type ChatRoom struct {
	Groups map[string][]string
}

func (c *ChatRoom) CreateGroup(group string) {
	c.Groups[group] = []string{}
}

func (c *ChatRoom) SendMessage(group string, message string) error {
	if _, exists := c.Groups[group]; !exists {
		return errors.New("group does not exist")
	}
	c.Groups[group] = append(c.Groups[group], message)
	return nil
}

func (c *ChatRoom) ReadMessage(group string) (string, error) {
	message, exists := c.Groups[group]
	if !exists {
		return "", errors.New("group does not exist")
	}

	if len(c.Groups[group]) == 0 {
		return "", errors.New("group message is empty")
	}

	msg := message[0]
	c.Groups[group] = message[1:]
	return msg, nil
}

func (c *ChatRoom) ListGroups() []string {
	var groups []string
	for i := range c.Groups {
		groups = append(groups, i)
	}
	return groups
}

func main() {
	chatRoom := &ChatRoom{
		Groups: make(map[string][]string),
	}

	chatRoom.CreateGroup("School")
	chatRoom.CreateGroup("Work")
	chatRoom.CreateGroup("Church")

	groups := chatRoom.ListGroups()
	for i, j := range groups {
		fmt.Printf("%d -> %s\n", i, j)
	}

	_ = chatRoom.SendMessage("School", "Three Assignment to be solved tonight")
	
	msg, _ := chatRoom.ReadMessage("School")
	fmt.Println(msg)
}