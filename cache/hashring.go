package cache

import (
	"hash/fnv"
)

type HashRing struct {
	nodes []*CacheNode
}

func NewHashRing(nodes []*CacheNode) *HashRing {
	return &HashRing{nodes: nodes}
}

func (h *HashRing) GetNode(key string) *CacheNode {
	hash := fnv.New32a()
	hash.Write([]byte(key))
	index := hash.Sum32() % uint32(len(h.nodes))
	return h.nodes[index]
}
