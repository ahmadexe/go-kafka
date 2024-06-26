# Go Kafka!
Kafka, developed by the Apache Software Foundation, is a high-throughput, distributed messaging system designed for building real-time data pipelines and streaming applications. It efficiently manages and processes large volumes of data in a publish-subscribe model, allowing producers to send messages to various topics and consumers to subscribe to those topics to receive the messages. Kafka's architecture ensures fault tolerance, scalability, and durability through features like partitioning, replication, and a distributed commit log.

## What is this repo about
This repo contains a basic example of how to use Kafka with Go. 

## Do you need Kafka?
Probably not, but nothing bad about learning it!

## How to run?
1. Start by cloning the repo
```
git clone git@github.com:ahmadexe/go-kafka.git
```

2. Navigate to the directory
3. Run the command in the root
   `docker-compose up -d`
4. Move to the producer directory and run
   `go run main.go`
6. Move to the consumer directory and run
   `go run main.go`
7. Pull up postman send a POST request at http://localhost:8080/message the body should look something like
   ```
    {
	"topic": "auth",
	"message": "Hi this is message 1"
    }
   ```

8. IF YOU CHANGE THE TOPIC CHANGE THAT IN THE CODE BASE AS WELL.
