package kafka

import (
	"fmt"
	"os"

	confluentKafka "github.com/confluentinc/confluent-kafka-go/kafka"
)

type KafkaConsumer struct {
	MsgChan chan *confluentKafka.Message
}

func NewKafkaConsumer(msgChan chan *confluentKafka.Message) *KafkaConsumer {
	return &KafkaConsumer{
		MsgChan: msgChan,
	}
}

func (k *KafkaConsumer) Consume() {
	configMap := &confluentKafka.ConfigMap{
		"bootstrap.servers": os.Getenv("KafkaBootstrapServers"),
		"group.id":          os.Getenv("KafkaConsumerGroupId"),
	}

	c, err := confluentKafka.NewConsumer(configMap)

	if err != nil {
		panic(err)
	}

	if err == nil {
		topics := []string{os.Getenv("KafkaReadTopic")}

		c.SubscribeTopics(topics, nil)

		fmt.Println("Kafka consumer has been started")

		for {
			msg, err := c.ReadMessage(-1)

			if err == nil {
				k.MsgChan <- msg
			}
		}
	}

}
