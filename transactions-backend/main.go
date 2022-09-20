package main

import (
	"fmt"
	"log"

	confluentKafka "github.com/confluentinc/confluent-kafka-go/kafka"
	kafka "github.com/felpssc/food-ordering-kafka-poc/transactions-backend/infra/kafka"
	"github.com/joho/godotenv"
)

func init() {
	err := godotenv.Load()

	fmt.Println("Loading .env file")

	if err != nil {
		log.Fatal("error loading .env file")
	}
}

func main() {
	msgChan := make(chan *confluentKafka.Message)
	consumer := kafka.NewKafkaConsumer(msgChan)

	consumer.Consume()

	for msg := range msgChan {
		// go Produce(msg)
		fmt.Println(string(msg.Value))
	}
}
