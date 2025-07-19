package kafka

import (
	"context"
	"fmt"
	"encoding/json"
	"github.com/segmentio/kafka-go"
)

// KafkaProducer struct to manage Kafka writer
type KafkaProducer struct {
	writer *kafka.Writer
}

// NewKafkaProducer creates a new Kafka producer with the specified broker address and topic
func NewKafkaProducer(brokerAddress string, topic string) *KafkaProducer {
	kafkaWriter := kafka.NewWriter(kafka.WriterConfig{
		Brokers: []string{brokerAddress},
		Topic:   topic,
		Balancer: &kafka.LeastBytes{},
	})

	return &KafkaProducer{
		writer: kafkaWriter,
		}
}

// Close closes the Kafka writer
func (kp *KafkaProducer) Close() error {
	if kp.writer != nil {
		return kp.writer.Close()
	}
	return nil
}

/// PublishStockReservedEvent publishes a stock reserved event to the Kafka topic
func (kp *KafkaProducer) PublishStockReservedEvent(ctx context.Context, reserve StockReserveEvent) error {
	value, err := json.Marshal(reserve)
	if err != nil {
		return fmt.Errorf("failed to marshal reserve event: %v", err)
	}
	event := kafka.Message{
		Key:   []byte(reserve.SKU),
		Value: value,
	}
	
	if err := kp.writer.WriteMessages(ctx, event); err != nil {
		return fmt.Errorf("failed to write message to Kafka: %v", err)
	}
	return nil
}

/// PublishStockReleasedEvent publishes a stock released event to the Kafka topic
func (kp *KafkaProducer) PublishStockReleasedEvent(ctx context.Context, release StockReleaseEvent) error {
	value, err := json.Marshal(release)
	if err != nil {
		return fmt.Errorf("failed to marshal release event: %v", err)
	}
	event := kafka.Message{
		Key:   []byte(release.SKU),
		Value: value,
	}
	if err := kp.writer.WriteMessages(ctx, event); err != nil {
		return fmt.Errorf("failed to write message to Kafka: %v", err)
	}
	return nil
}