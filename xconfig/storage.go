package xconfig

import (
	"context"
	"sync"

	"github.com/go-redis/redis/v8"
	clientv3 "go.etcd.io/etcd/client/v3"
)

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

type RedisStorage struct {
	conn *redis.Client
	ctx  context.Context
}

func NewRedisStorage(addr string) *RedisStorage {
	return &RedisStorage{
		conn: redis.NewClient(&redis.Options{
			Addr: addr,
		}),
		ctx: context.Background(),
	}
}

func (r *RedisStorage) Set(key string, value []byte) error {
	return r.conn.Set(r.ctx, key, value, 0).Err()
}

func (r *RedisStorage) Get(key string) ([]byte, error) {
	val, err := r.conn.Get(r.ctx, key).Bytes()
	if err != nil {
		if err == redis.Nil {
			return nil, ErrNotFound
		}
		return nil, err
	}

	if len(val) == 0 || val == nil {
		return nil, ErrNotFound
	}

	return val, nil
}

func (r *RedisStorage) Del(key string) error {
	return r.conn.Del(r.ctx, key).Err()
}

func (r *RedisStorage) Keys() ([]string, error) {
	return r.conn.Keys(r.ctx, "*").Result()
}

var _ Storage = (*LocalStorage)(nil)

type EtcdStorage struct {
	ctx  context.Context
	conn *clientv3.Client
}

func NewEtcdStorage(endpoints []string) *EtcdStorage {
	conn, err := clientv3.New(clientv3.Config{
		Endpoints: endpoints,
	})
	if err != nil {
		panic(err)
	}
	return &EtcdStorage{
		ctx:  context.Background(),
		conn: conn,
	}
}

func (e *EtcdStorage) Set(key string, value []byte) error {
	_, err := e.conn.Put(e.ctx, key, string(value))
	if err != nil {
		return err
	}
	return nil
}

func (e *EtcdStorage) Get(key string) ([]byte, error) {
	resp, err := e.conn.Get(e.ctx, key)
	if err != nil {
		return nil, err
	}
	if len(resp.Kvs) == 0 {
		return nil, ErrNotFound
	}
	return resp.Kvs[0].Value, nil
}

func (e *EtcdStorage) Del(key string) error {
	_, err := e.conn.Delete(e.ctx, key)
	if err != nil {
		return err
	}
	return nil
}
