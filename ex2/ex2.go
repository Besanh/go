package ex2

import (
	"errors"
	"fmt"
	"math/rand"
	"time"
)

func Ex2(name string) (string, error) {
	if name == "" {
		return "error roi", errors.New("Co loi roi")
	}
	// message := fmt.Sprintf("Wao %v", name)
	message := fmt.Sprintf(randomName(), name)
	return message, nil
}

func Ex3(names []string) (map[string]string, error) {
	messages := make(map[string]string, 0)
	for _, name := range names {
		message, error := Ex2(name)
		if error != nil {
			return nil, error
		}
		messages[name] = message
	}
	return messages, nil
}

func init() {
	rand.Seed(time.Now().UnixNano())
}

func randomName() string {
	arr := []string{
		"Hi %v",
		"Hello %v",
		"AAAA %v",
	}

	return arr[rand.Intn(len(arr))]
}
