package main

import (
	"encoding/json"
	"math/rand"
	"time"

	confluentKafka "github.com/confluentinc/confluent-kafka-go/kafka"
	order "github.com/felpssc/food-ordering-kafka-poc/transactions-backend/app/order"
	kafka "github.com/felpssc/food-ordering-kafka-poc/transactions-backend/infra/kafka"
)

func Produce(msg *confluentKafka.Message) {
	producer := kafka.NewKafkaProducer()

	objectOrder := NewOrder()

	json.Unmarshal(msg.Value, &objectOrder)

	// choose an random status between "confirmed" and "cancelled"
	status := []string{"confirmed", "cancelled"}

	objectOrder.status = status[RandomNumber(0, 1)]

	// sleep for 2.5 seconds
	time.Sleep(2500 * time.Millisecond)

	parsedJson, _ := json.Marshal(objectOrder)

	// publish the order to the topic
	order.Publish(string(parsedJson), "order_confirmed", producer)
}

func RandomNumber(min, max int) int {
	return rand.Intn(max-min) + min
}
