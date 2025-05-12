package pkg

import (
	"context"
	"fmt"
	"sync"

	"github.com/segmentio/kafka-go"
)

// var onceDo sync.Once
// var kafkaClient *kafka.Writer
// var kafkaClientDriver *kafka.Conn

var (
	kafkaWriters = map[string]*kafka.Writer{}
	kafkaMu      sync.Mutex
)

func KafkaConnector(topic string) *kafka.Writer {
	kafkaMu.Lock()
	defer kafkaMu.Unlock()

	if w, ok := kafkaWriters[topic]; ok {
		return w
	}

	writer := &kafka.Writer{
		Addr:     kafka.TCP("localhost:9092"),
		Topic:    topic,
		Balancer: &kafka.LeastBytes{},
	}

	kafkaWriters[topic] = writer
	return writer
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
		fmt.Printf("msg_order=%s\n", string(m.Value))
	}
}

func KafkaReadAssign() {
	kfkCnsmr := kafka.NewReader(kafka.ReaderConfig{
		Brokers:     []string{"localhost:9092"},
		Topic:       "order-assigned",
		StartOffset: kafka.FirstOffset,
		MaxBytes:    10e6,
	})

	for {
		m, err := kfkCnsmr.ReadMessage(context.Background())
		if err != nil {
			continue
		}
		fmt.Printf("msg_assign=%s\n", string(m.Value))
	}
}
