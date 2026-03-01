package kafka

import (
	"log"
	"os"
	"strconv"
	"strings"
)

type KafkaConfig struct {
	Brokers            []string
	DefaultPartitions  int
	DefaultReplication int
}

var kafkaCfg KafkaConfig

func init() {
	kc := KafkaConfig{}
	kc.Brokers = strings.Split(getEnv("KAFKA_BROKERS", "localhost:9092"), ",")
	dfltPart := getEnv("KAFKA_DEFAULT_PARTITIONS", "3")
	dfltRepl := getEnv("KAFKA_REPLICATION_FACTOR", "1")

	if i, err := strconv.Atoi(dfltPart); err != nil {
		log.Println("Error in converting KAFKA_DEFAULT_PARTITIONS to int:", err)
		log.Println("Defaulting to 3")
	} else {
		kc.DefaultPartitions = i
	}

	if i, err := strconv.Atoi(dfltRepl); err != nil {
		log.Println("Error in converting KAFKA_REPLICATION_FACTOR to int:", err)
		log.Println("Defaulting to 1")
	} else {
		kc.DefaultReplication = i
	}

	kafkaCfg = kc
	log.Println("KafkaConfig Loaded:", kafkaCfg)
}

func getEnv(key, fallback string) string {
	if value, exists := os.LookupEnv(key); exists && len(value) > 0 {
		log.Printf("Found env var %s=%s", key, value)
		return value
	}
	log.Printf("Env var %s not set, using default=%s", key, fallback)
	return fallback
}
