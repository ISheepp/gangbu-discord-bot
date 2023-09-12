package queue

import (
	"crypto/tls"
	"gangbu/pkg/util"
	"github.com/segmentio/kafka-go"
	"github.com/segmentio/kafka-go/sasl/scram"
	"os"
)

func NewKafkaWriter() (*kafka.Writer, error) {
	mechanism, err := scram.Mechanism(scram.SHA256, os.Getenv("KAFKA_USERNAME"), os.Getenv("KAFKA_PASSWORD"))
	if err != nil {
		util.Logger.Error("创建kafka writer失败", err)
		return nil, err
	}

	sharedTransport := &kafka.Transport{
		SASL: mechanism,
		TLS:  &tls.Config{},
	}

	w := &kafka.Writer{
		Addr:      kafka.TCP(os.Getenv("KAFKA_URL")),
		Topic:     "game",
		Balancer:  &kafka.Hash{},
		Transport: sharedTransport,
		Logger:    util.Logger,
	}
	return w, nil

}

func NewKafkaReader() (*kafka.Reader, error) {
	mechanism, err := scram.Mechanism(scram.SHA256, os.Getenv("KAFKA_USERNAME"), os.Getenv("KAFKA_PASSWORD"))
	if err != nil {
		util.Logger.Error("创建kafka reader失败", err)
		return nil, err
	}

	dialer := &kafka.Dialer{
		TLS:           &tls.Config{},
		SASLMechanism: mechanism,
	}

	r := kafka.NewReader(kafka.ReaderConfig{
		Brokers:     []string{os.Getenv("KAFKA_URL")},
		GroupID:     "test-1",
		Topic:       "game",
		Dialer:      dialer,
		ErrorLogger: util.Logger,
	})
	return r, nil
}
