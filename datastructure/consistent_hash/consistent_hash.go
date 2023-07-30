package consistenthash

import (
	"fmt"
	"hash/crc32"
	"sort"
)

type (
	HashFn func(data []byte) uint32

	ConsistentHash struct {
		Hash     HashFn
		Replicas int
		Ring     []int
		Map      map[int]string
	}
)

func New(replicas int, fn HashFn) *ConsistentHash {
	m := &ConsistentHash{
		Replicas: replicas,
		Hash:     fn,
		Map:      make(map[int]string),
	}
	if m.Hash == nil {
		m.Hash = crc32.ChecksumIEEE
	}
	return m
}

func (c *ConsistentHash) Add(nodes ...string) {
	for _, node := range nodes {
		for i := 0; i < c.Replicas; i++ {
			hash := int(c.Hash([]byte(fmt.Sprintf("%s-%d", node, i))))
			c.Ring = append(c.Ring, hash)
			c.Map[hash] = node
		}
	}
	sort.Ints(c.Ring)
}

func (c *ConsistentHash) Get(key string) string {
	if len(c.Ring) == 0 {
		return ""
	}
	hash := int(c.Hash([]byte(key)))
	idx := sort.Search(len(c.Ring), func(i int) bool {
		return c.Ring[i] >= hash
	})
	return c.Map[c.Ring[idx%len(c.Ring)]]
}

func (c *ConsistentHash) Remove(key string) bool {
	for i := 0; i < c.Replicas; i++ {
		hash := int(c.Hash([]byte(fmt.Sprintf("%s-%d", key, i))))
		if _, ok := c.Map[hash]; !ok {
			return false
		}
		delete(c.Map, hash)
		for j, v := range c.Ring {
			if v == hash {
				c.Ring = append(c.Ring[:j], c.Ring[j+1:]...)
			}
		}
	}
	sort.Ints(c.Ring)
	return true
}
