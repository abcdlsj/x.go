package producer

import (
	"context"
	"testing"

	kafka "github.com/segmentio/kafka-go"
)

func TestProducer(t *testing.T) {
	var p = NewProducer()

	w := p.NewWriter()

	err := w.WriteMessages(
		context.Background(),
		kafka.Message{Value: []byte("one!")},
		kafka.Message{Value: []byte("two!")},
		kafka.Message{Value: []byte("three!")},
	)
	if err != nil {
		t.Fatalf("could not write messages: %v", err)
	}

	w.Close()
}
