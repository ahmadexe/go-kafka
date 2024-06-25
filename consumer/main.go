package main

import (
	"context"
	"encoding/json"
	"fmt"
	"log"
	"time"

	"github.com/ahmadexe/go-kafka/data"
	"github.com/segmentio/kafka-go"
)

const (
	topic = "auth"
	partition = 0
	offset = kafka.LastOffset
)

func main() {
	r := kafka.NewReader(kafka.ReaderConfig{
		Brokers:   []string{"localhost:29092"},
		Topic:     topic,
		Partition: partition,
		MaxBytes:  10e6, // 10MB
		CommitInterval: time.Second,
	})

	r.SetOffset(offset)

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
