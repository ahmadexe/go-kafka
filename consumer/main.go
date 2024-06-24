package main

import (
	"context"
	"encoding/json"
	"fmt"
	"log"

	"github.com/ahmadexe/go-kafka/data"
	"github.com/segmentio/kafka-go"
)

func main() {
	topic := "auth"
	partition := 0

	r := kafka.NewReader(kafka.ReaderConfig{
		Brokers:   []string{"localhost:29092"},
		Topic:     topic,
		Partition: partition,
		MaxBytes:  10e6, // 10MB
	})

	for {
		m, err := r.ReadMessage(context.Background())
		if err != nil {
			break
		}
		var message data.Message
		if err := json.Unmarshal(m.Value, &message); err != nil {
			log.Fatal("failed to unmarshal message:", err)
		}

		fmt.Printf("message: %s\n", message.Message)
	}
	
	if err := r.Close(); err != nil {
		log.Fatal("failed to close reader:", err)
	}
}
