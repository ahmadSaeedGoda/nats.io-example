package main

import (
	"bufio"
	"flag"
	"fmt"
	"log"
	"os"
	"strings"
	"time"

	"github.com/nats-io/nats.go"
)

const (
	natsURL = "nats://localhost:4222"
	subject = "chat"
)

func main() {
	// Parse the username from command-line arguments
	username := flag.String("username", "", "Your chat username")
	flag.Parse()

	if *username == "" {
		log.Fatalf("Username is required. Usage: go run chat.go -username=<your_username>")
	}

	// Connect to NATS
	nc, err := nats.Connect(natsURL)
	if err != nil {
		log.Fatal(err)
	}
	defer nc.Close()

	// Subscribe to the chat subject
	_, err = nc.Subscribe(subject, func(m *nats.Msg) {
		fmt.Println(string(m.Data))
	})
	if err != nil {
		log.Fatal(err)
	}

	// Use a buffered channel to process messages
	ch := make(chan string)
	go func() {
		scanner := bufio.NewScanner(os.Stdin)
		for scanner.Scan() {
			ch <- scanner.Text()
		}
	}()

	fmt.Println("Welcome to the chat! Type your messages below:")

	for msg := range ch {
		if strings.TrimSpace(msg) == "" {
			log.Println("Empty message!", "not sent.")
			continue
		}
		chatMessage := fmt.Sprintf("[%s] %s: %s", time.Now().Format("15:04:05"), *username, msg)
		nc.Publish(subject, []byte(chatMessage))
	}
}
