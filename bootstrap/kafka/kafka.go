package kafka

import (
	"backend-task/bootstrap/config"
	"backend-task/pkg/logger"
	"encoding/json"
	"fmt"
	"github.com/confluentinc/confluent-kafka-go/kafka"
	"github.com/google/uuid"
	"go.uber.org/fx"
)

type Message struct {
	Body       interface{} `json:"body"`
	Properties []int       `json:"properties"`
	Headers    struct {
		contentType   string `json:"contentType"`
		correlationID string `json:"correlationID"`
	} `json:"headers"`
}

type Producer interface {
	SendMessage(topic string, body interface{}) error
}

type producer struct {
	producer *kafka.Producer
}

func NewKafkaProducer(config *config.Config) Producer {
	kafkaBroker := fmt.Sprintf("%s:%s", config.KafkaHost, config.KafkaPort)
	logger.Log.Info("Connecting to kafka broker: ", kafkaBroker)
	kafkaProducer, err := kafka.NewProducer(
		&kafka.ConfigMap{
			"bootstrap.servers":     kafkaBroker,
			"broker.address.family": "v4",
			"security.protocol":     "PLAINTEXT",
		},
	)
	if err != nil {
		logger.Log.Error(err)
		return nil
	}

	logger.Log.Info("Successfully connected to kafka broker...")

	go func() {
		for e := range kafkaProducer.Events() {
			switch ev := e.(type) {
			case *kafka.Message:
				if ev.TopicPartition.Error != nil {
					fmt.Printf("Delivery failed: %v\n", ev.TopicPartition.Error)

				} else {
					fmt.Printf(
						"Delivered message to topic %s [%d] at offset %v\n",
						*ev.TopicPartition.Topic, ev.TopicPartition.Partition, ev.TopicPartition.Offset,
					)
				}
			case kafka.Error:
				fmt.Printf("Error: %v\n", ev)

			default:
				fmt.Printf("Ignored event: %s\n", ev)
			}

		}
	}()
	return &producer{kafkaProducer}
}

func (p producer) SendMessage(topic string, body interface{}) error {
	correlationID := uuid.New().String()

	data := Message{
		Body:       body,
		Properties: []int{},
	}
	data.Headers.contentType = "application/json"
	data.Headers.correlationID = correlationID

	source, err := json.Marshal(data)
	if err != nil {
		logger.Log.Warn("Error to parser data to byte", err)
		return err
	}

	err = p.producer.Produce(
		&kafka.Message{
			TopicPartition: kafka.TopicPartition{Topic: &topic, Partition: kafka.PartitionAny},
			Value:          source,
			Key:            []byte(correlationID),
			Headers: []kafka.Header{
				{Key: "correlationId", Value: []byte(correlationID)},
				{Key: "applicationName", Value: []byte("ms-miniapp")},
			},
		}, nil,
	)
	if err != nil {
		logger.Log.Warn("Error to send message to kafka producer", err)
		return err
	}

	p.producer.Flush(-1)
	logger.Log.Info("Message sent to kafka producer")

	return nil
}

var Module = fx.Module(
	"client",
	fx.Provide(NewKafkaProducer),
)
