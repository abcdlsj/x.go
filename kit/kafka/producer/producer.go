package producer

import (
	"crypto/tls"
	"log"
	"os"
	"time"

	kafka "github.com/segmentio/kafka-go"
	"github.com/segmentio/kafka-go/sasl/scram"
)

type Producer struct {
	conn *kafka.Dialer
}

func NewProducer() *Producer {
	mechanism, err := scram.Mechanism(scram.SHA256, env("KAFKA_USERNAME"), env("KAFKA_PASSWORD"))
	if err != nil {
		log.Fatalln(err)
	}

	dialer := &kafka.Dialer{
		SASLMechanism: mechanism,
		TLS:           &tls.Config{},
		Timeout:       3 * time.Second,
	}

	return &Producer{
		conn: dialer,
	}
}

func (p *Producer) NewWriter() *kafka.Writer {
	return kafka.NewWriter(kafka.WriterConfig{
		Brokers: []string{env("KAFKA_ENDPOINT")},
		Topic:   env("KAFKA_TOPIC"),
		Dialer:  p.conn,
	})
}

func env(k string) string {
	return os.Getenv(k)
}
