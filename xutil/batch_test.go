package xutil

import "testing"

func TestBatch(t *testing.T) {
	batch := NewBatch(0, 100, 12)
	for batch.HasNext() {
		t.Log(batch.Next())
	}

	// /usr/local/bin/go test -v -timeout 30s -run ^TestBatch$ github.com/abcdlsj/x/xutil
}
