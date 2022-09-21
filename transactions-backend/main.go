package main

import (
	"fmt"
	"log"

	confluentKafka "github.com/confluentinc/confluent-kafka-go/kafka"
	"github.com/felpssc/food-ordering-kafka-poc/app/order"
	kafka "github.com/felpssc/food-ordering-kafka-poc/infra/kafka"
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

	go consumer.Consume()

	for msg := range msgChan {
		go order.Produce(msg)
	}
}
