package kafka

import (
	"context"
	"log"
	"strings"
	"time"

	"github.com/segmentio/kafka-go"
)

var writer *kafka.Writer

var ctxBg = context.Background()

type KafkaMsg = kafka.Message

func NewWriter() *kafka.Writer {
	// if err := EnsureTopicExists(topic); err != nil {
	// 	return nil
	// }

	return &kafka.Writer{
		Addr:                   kafka.TCP(kafkaCfg.Brokers...),
		Balancer:               &kafka.LeastBytes{},
		AllowAutoTopicCreation: true,
		// Topic:    topic, // should be defined per message
	}
}

func NewReader(topic string) *kafka.Reader {
	if err := EnsureTopicExists(topic); err != nil {
		return nil
	}

	return kafka.NewReader(kafka.ReaderConfig{
		Brokers:          kafkaCfg.Brokers,
		Topic:            topic,
		GroupID:          topic + "-readers",
		ReadBatchTimeout: 1,
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

func PublishToTopic(data []byte, topic string) {
	if writer == nil {
		writer = NewWriter()
	}

	ctx, cancel := context.WithTimeout(ctxBg, 2*time.Second)
	defer cancel()
	kfkMsg := kafka.Message{
		Topic: topic,
		Value: data,
	}
	if err := writer.WriteMessages(ctx, kfkMsg); err != nil {
		log.Println("There was an error in Pushing message", kfkMsg, "to Topic", topic, ".")
	}
}

func ConsumeFromTopic(topic string, topicChan chan<- kafka.Message) {
	reader := NewReader(topic)
	log.Println("reader", reader)
	defer reader.Close()
	for {
		if msg, err := reader.ReadMessage(ctxBg); err != nil {
			log.Println("There was some error in Reading message from Topic", topic, "", err)
		} else {
			log.Println("Message from Topic", topic, "", string(msg.Value))
			topicChan <- msg
		}
	}
}
