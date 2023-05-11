package consumer

import (
	"context"
	"testing"
)

func TestConsumer(t *testing.T) {
	var c = NewConsumer()

	r := c.NewReader()
	for {
		m, err := r.ReadMessage(context.Background())
		if err != nil {
			break
		}
		t.Logf("message at offset %d: %s = %s\n", m.Offset, string(m.Key), string(m.Value))
	}

	r.Close()
}
