package kafka

import (
	ckafka "github.com/confluentinc/confluent-kafka-go/kafka"
	"os"
)

type KafkaConsumer struct {
	MsgChan chan *ckafka.Message
}

func(k *KafkaConsumer) Consume() {
	configMap := &ckafka.ConfigMap{
		"bootstrap.servers": os.Getenv("KafkaBootstrapServers"),
		"group.id": os.Getenv("KafkaConsumerGroupId"),
	}
	c, err := ckafka.NewConsumer(configMap)
	if err != nil {
		log.Fatal("error consuming kafka message:" + err.Error())
	}
	topics := []string{os.Getenv("KafkaReadTopic")}
	c.SubscribeTopics(topics, nil)
	fmt.Println("Kafka consumer has been started")
	for {
		msg, err := c.ReadMessage(-1)
		if err == nill {
			k.MsgChan <- msg
		}
	}
}

