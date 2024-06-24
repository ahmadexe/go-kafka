package main

import (
	"context"
	"encoding/json"
	"log"

	"github.com/ahmadexe/go-kafka/data"
	"github.com/gin-gonic/gin"
	"github.com/segmentio/kafka-go"
)

func main() {
	router := gin.Default()
	router.POST("/message", pushMessageToQueue)
	router.Run(":8080")
}

func pushMessageToQueue(ctx *gin.Context) {
	var message data.Message

	if err := ctx.BindJSON(&message); err != nil {
		ctx.JSON(400, gin.H{"error": err.Error()})
		return
	}

	byteMessage, err := json.Marshal(message)
	if err != nil {
		log.Fatal("failed to marshal message:", err)
	}

	partition := 0
	conn, err := kafka.DialLeader(context.Background(), "tcp", "localhost:29092", message.Topic, partition)

	if err != nil {
		log.Fatal("failed to dial leader:", err)
	}
	defer conn.Close()

	_, err = conn.WriteMessages(
		kafka.Message{Value: byteMessage},
	)
	if err != nil {
		log.Fatal("failed to write messages:", err)
	}

	ctx.JSON(200, gin.H{"message": "success", "data": message})
}
