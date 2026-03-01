package kafka

import (
	"log"
	"strings"

	"github.com/segmentio/kafka-go"
)

func NewProducer(topic string) *kafka.Writer {
	if err := EnsureTopicExists(topic); err != nil {
		return nil
	}

	return &kafka.Writer{
		Addr:     kafka.TCP(kafkaCfg.Brokers...),
		Topic:    topic,
		Balancer: &kafka.LeastBytes{},
	}
}

func NewConsumer(topic string) *kafka.Reader {
	if err := EnsureTopicExists(topic); err != nil {
		return nil
	}

	return kafka.NewReader(kafka.ReaderConfig{
		Brokers: kafkaCfg.Brokers,
		Topic:   topic,
		GroupID: topic + "-consumer",
	})
}

func CreateTopic(name string) error {
	if cont, err := GetKafkaController(); err != nil {
		log.Println("Error in dialing Kafka Controller:", err)
		return err
	} else {
		defer cont.Close()
		err := cont.CreateTopics(kafka.TopicConfig{
			Topic:             name,
			NumPartitions:     kafkaCfg.DefaultPartitions,
			ReplicationFactor: kafkaCfg.DefaultReplication,
		})

		if err != nil {
			log.Println("Error in Creating Kafka Topic", name, ", Error:", err)
		}
		return err
	}
}

func EnsureTopicExists(topic string) error {
	if err := CreateTopic(topic); err != nil {
		if strings.Contains(err.Error(), "Topic already exists") {
			log.Println("Topic", topic, "already exists")
		} else {
			log.Println("Error in Ensuring Kafka Topic Exists:", err)
			return err
		}
	}

	log.Println("Topic", topic, "ensured.")
	return nil
}
