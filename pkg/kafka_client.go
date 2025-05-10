package pkg

import (
	"context"
	"fmt"
	"sync"

	"github.com/segmentio/kafka-go"
)

var onceDo sync.Once
var kafkaClient *kafka.Conn

func KafkaConnector() *kafka.Conn {
	onceDo.Do(func() {
		conn, err := kafka.DialLeader(context.Background(), "tcp", "localhost:9092", "order-created", 0)

		if err != nil {
			panic("err when connect kafka")
		}

		fmt.Print("Kafka is initialized :)")
		kafkaClient = conn

	})
	return kafkaClient
}

func KafkaReadOrderCreated() {
	kfkCnsmr := kafka.NewReader(kafka.ReaderConfig{
		Brokers:     []string{"localhost:9092"},
		Topic:       "order-created",
		StartOffset: kafka.FirstOffset,
		MaxBytes:    10e6,
	})

	for {
		m, err := kfkCnsmr.ReadMessage(context.Background())
		if err != nil {
			continue
		}
		fmt.Printf("message=%s\n", string(m.Value))
	}

}
