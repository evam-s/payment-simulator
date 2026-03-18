package kafka

import (
	"fmt"
	"log"

	"github.com/segmentio/kafka-go"
)

func GetKafkaController() (*kafka.Conn, error) {
	for _, broker := range kafkaCfg.Brokers {
		brokerConn, err := kafka.Dial("tcp", broker)
		if err != nil {
			log.Println("Error in Connecting to Broker ", broker, ", Error:", err)
			// brokerConn.Close()
			continue
		}

		contUri, err1 := brokerConn.Controller()
		brokerConn.Close()
		if err1 != nil {
			log.Println("Error in Getting the Controller Data:", err1)
			continue
		}

		contAddr := fmt.Sprintf("%s:%d", contUri.Host, contUri.Port)
		cont, err2 := kafka.Dial("tcp", contAddr)
		if err2 != nil {
			log.Println("Error in Connecting to Broker ", broker, ", Error:", err)
			continue
		}

		return cont, nil // Controller connection must be closed by called func.
	}
	panic(fmt.Errorf("Unable to connect to any of the provided Brokers."))
}
