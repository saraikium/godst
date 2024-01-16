package maps

import (
	"cmp"
	"fmt"
	"hash/fnv"
)

type HashTable[V any] struct {
	buckets [][]KeyValue[V]
	size    int
}

type KeyValue[V any] struct {
	key   string
	value V
}

func New[V any](size int) *HashTable[V] {
	return &HashTable[V]{
		buckets: make([][]KeyValue[V], 10),
		size:    size,
	}
}

func (h *HashTable[V]) Put(key string, value V) {
	index := h.hash(key)
	h.buckets[index] = append(h.buckets[index], KeyValue[V]{key: key, value: value})

}

func (h *HashTable[V]) hash(key string) int {
	hasher := fnv.New32a()
	hasher.Write([]byte(key))
	return int(hasher.Sum32()) % h.size
}

func (h *HashTable[V]) Get(key string) (V, error) {
	index := h.hash(key)
	bucket := h.buckets[index]

	for _, kv := range bucket {
		if kv.key == key {
			return kv.value, nil
		}
	}

	return *new(V), fmt.Errorf("Key not found.")

}
