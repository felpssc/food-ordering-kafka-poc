package order

import (
	"encoding/json"
	"fmt"
	"math/rand"
	"time"

	confluentKafka "github.com/confluentinc/confluent-kafka-go/kafka"
	kafka "github.com/felpssc/food-ordering-kafka-poc/infra/kafka"
)

func Produce(msg *confluentKafka.Message) {
	producer := kafka.NewKafkaProducer()

	var order Order

	err := json.Unmarshal(msg.Value, &order)

	if err != nil {
		fmt.Println("Error while unmarshalling order: ", err.Error())
	}

	transaction_id := generateTransactionId()

	order.Transaction_id = transaction_id

	// random status 0 or 1
	status := RandomNumber(0, 2)

	statusEnum := []string{"confirmed", "cancelled"}

	order.Status = statusEnum[status]

	orderJson, err := json.Marshal(order)

	if err != nil {
		fmt.Println("Error while marshalling order: ", err.Error())
	}

	queueEnum := []string{"order_confirmed", "order_cancelled"}

	// sleep for 1.5 seconds
	time.Sleep(1500 * time.Millisecond)

	kafka.Publish(string(orderJson), queueEnum[status], producer)

	fmt.Println(
		"Order: ", string(order.Document_id),
		" sent to queue: ", queueEnum[status],
		" with transaction_id: ", transaction_id)
}

func RandomNumber(min, max int) int {
	return rand.Intn(max-min) + min
}

func generateTransactionId() string {
	return fmt.Sprintf("%d", time.Now().UnixNano())
}
