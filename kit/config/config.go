package config

import "sync"

type Service interface {
	Get(key string) ([]byte, error)
	Set(key string, value []byte) error
	Del(key string) error
	Watch(key string) (chan Event, error)
}

type Event struct {
	Type EventType
	Key  string
	Val  []byte
}

type Client struct {
	storage  Storage
	watchers map[string]chan Event
	mu       sync.RWMutex
}

type Option func(*Client)

func WithStorage(s Storage) Option {
	return func(c *Client) {
		c.storage = s
	}
}

func NewClient(options ...Option) *Client {
	c := &Client{
		watchers: make(map[string]chan Event),
	}
	for _, opt := range options {
		opt(c)
	}
	return c
}

func NewLocalClient() *Client {
	return NewClient(WithStorage(NewLocalStorage()))
}

func (l *Client) Get(key string) ([]byte, error) {
	l.mu.RLock()
	defer l.mu.RUnlock()
	return l.storage.Get(key)
}

func (l *Client) Set(key string, val []byte) error {
	l.mu.Lock()
	defer l.mu.Unlock()

	et := Unknown
	if _, err := l.storage.Get(key); err != nil {
		et = Create
	} else {
		et = Update
	}
	if err := l.storage.Set(key, val); err != nil {
		return err
	}
	if _, ok := l.watchers[key]; ok {
		l.watchers[key] <- Event{et, key, val}
	}
	return nil
}

func (l *Client) Del(key string) error {
	l.mu.Lock()
	defer l.mu.Unlock()
	if err := l.storage.Del(key); err != nil {
		return err
	}
	if _, ok := l.watchers[key]; ok {
		l.watchers[key] <- Event{Delete, key, nil}
	}
	return nil
}

func (l *Client) Watch(key string) (chan Event, error) {
	l.mu.Lock()
	defer l.mu.Unlock()
	if _, ok := l.watchers[key]; !ok {
		l.watchers[key] = make(chan Event, 1)
	}
	return l.watchers[key], nil
}
