package config

import "testing"

var client = NewLocalClient()

func TestCRUD(t *testing.T) {
	if err := client.Set("foo", []byte("bar")); err != nil {
		t.Fatal(err)
	}
	val, err := client.Get("foo")
	if err != nil {
		t.Fatal(err)
	}
	if string(val) != "bar" {
		t.Fatalf("expected %s, got %s", "bar", string(val))
	}

	if err := client.Del("foo"); err != nil {
		t.Fatal(err)
	}

	if _, err := client.Get("foo"); err != ErrNotFound {
		t.Fatalf("expected %v, got %v", ErrNotFound, err)
	}

	if err := client.Set("foo", []byte("bar1")); err != nil {
		t.Fatal(err)
	}

	if err := client.Set("foo", []byte("bar2")); err != nil {
		t.Fatal(err)
	}

	val, err = client.Get("foo")
	if err != nil {
		t.Fatal(err)
	}
	if string(val) != "bar2" {
		t.Fatalf("expected %s, got %s", "bar2", string(val))
	}

	watcher, err := client.Watch("foo")
	if err != nil {
		t.Fatal(err)
	}

	if err := client.Set("foo", []byte("bar3")); err != nil {
		t.Fatal(err)
	}

	event := <-watcher
	if event.Type != Update {
		t.Fatalf("expected %v, got %v", Update, event.Type)
	}

	if string(event.Val) != "bar3" {
		t.Fatalf("expected %s, got %s", "bar3", string(event.Val))
	}

	if err := client.Del("foo"); err != nil {
		t.Fatal(err)
	}

	event = <-watcher
	if event.Type != Delete {
		t.Fatalf("expected %v, got %v", Delete, event.Type)
	}
}
