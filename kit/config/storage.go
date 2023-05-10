package config

import "sync"

type Storage interface {
	Get(key string) ([]byte, error)
	Set(key string, value []byte) error
	Del(key string) error
}

type LocalStorage struct {
	data map[string][]byte
	mu   sync.RWMutex
}

func NewLocalStorage() Storage {
	return &LocalStorage{
		data: make(map[string][]byte),
	}
}

func (l *LocalStorage) Get(key string) ([]byte, error) {
	l.mu.RLock()
	defer l.mu.RUnlock()
	if val, ok := l.data[key]; ok {
		return val, nil
	}
	return nil, ErrNotFound
}

func (l *LocalStorage) Set(key string, val []byte) error {
	l.mu.Lock()
	defer l.mu.Unlock()
	l.data[key] = val
	return nil
}

func (l *LocalStorage) Del(key string) error {
	l.mu.Lock()
	defer l.mu.Unlock()
	if _, ok := l.data[key]; !ok {
		return ErrNotFound
	}
	delete(l.data, key)
	return nil
}

type RedisClient struct{}
type EtcdClient struct{}
type S3Client struct{}
