package cache

import (
	"sync"
)

type CacheNode struct {
	cache map[string]string
	mutex sync.RWMutex
}

func NewCacheNode() *CacheNode {
	return &CacheNode{
		cache: make(map[string]string),
	}
}

func (n *CacheNode) Get(key string) (string, error) {
	n.mutex.RLock()
	defer n.mutex.RUnlock()
	value, ok := n.cache[key]
	if !ok {
		return "", nil
	}
	return value, nil
}

func (n *CacheNode) Set(key string, value string) error {
	n.mutex.Lock()
	defer n.mutex.Unlock()
	n.cache[key] = value
	return nil
}

func (n *CacheNode) Delete(key string) error {
	n.mutex.Lock()
	defer n.mutex.Unlock()
	delete(n.cache, key)
	return nil
}

type DistributedCache struct {
	nodes    []*CacheNode
	hashRing *HashRing
}

func NewDistributedCache() *DistributedCache {
	nodes := []*CacheNode{
		NewCacheNode(),
		NewCacheNode(),
		NewCacheNode(),
	}
	hashRing := NewHashRing(nodes)
	return &DistributedCache{
		nodes:    nodes,
		hashRing: hashRing,
	}
}

func (d *DistributedCache) Get(key string) (string, error) {
	node := d.hashRing.GetNode(key)
	return node.Get(key)
}

func (d *DistributedCache) Set(key string, value string) error {
	node := d.hashRing.GetNode(key)
	return node.Set(key, value)
}

func (d *DistributedCache) Delete(key string) error {
	node := d.hashRing.GetNode(key)
	return node.Delete(key)
}
