// implementation of hash map. how to hash. CRUD operations
package src

import (
	"fmt"
	"hash/fnv"
	"sync"
)

type BucketNode struct {
	Key   string
	Value any
	Next  *BucketNode
}

type Bucket struct {
	head *BucketNode
}

func NewBucket() *Bucket {
	return &Bucket{}
}

func NewBucketList(size int) []*Bucket {
	buckets := make([]*Bucket, size)
	for i := range buckets {
		buckets[i] = NewBucket()
	}
	return buckets
}

type HashMap struct {
	buckets    []*Bucket
	size       int
	capacity   int
	loadFactor float64
	mu         sync.RWMutex
}

func NewHashMap(size int) *HashMap {
	return &HashMap{
		buckets:    NewBucketList(size),
		capacity:   size,
		loadFactor: 0.75,
	}
}

func HashFunction(key string, size int) int {
	hash := fnv.New32a()
	hash.Write([]byte(key))
	return int(hash.Sum32()) % size
}

// Set inserts or updates a key-value pair
func (h *HashMap) Set(key string, value any) {
	h.mu.Lock()
	defer h.mu.Unlock()

	// Resize if load factor exceeded
	if float64(h.size)/float64(h.capacity) > h.loadFactor {
		h.resize()
	}

	index := HashFunction(key, h.capacity)
	bucket := h.buckets[index]

	// Check if key exists, update it
	current := bucket.head
	for current != nil {
		if current.Key == key {
			current.Value = value
			return
		}
		current = current.Next
	}

	// Insert at head (chaining)
	newNode := &BucketNode{Key: key, Value: value, Next: bucket.head}
	bucket.head = newNode
	h.size++
}

func (h *HashMap) Get(key string) (any, bool) {
	h.mu.RLock()
	defer h.mu.RUnlock()

	index := HashFunction(key, h.capacity)
	bucket := h.buckets[index]

	current := bucket.head
	for current != nil {
		if current.Key == key {
			return current.Value, true
		}
		current = current.Next
	}
	return nil, false
}

func (h *HashMap) Delete(key string) {
	h.mu.Lock()
	defer h.mu.Unlock()

	index := HashFunction(key, h.capacity)
	bucket := h.buckets[index]

	var prev *BucketNode
	current := bucket.head
	for current != nil {
		if current.Key == key {
			if prev == nil {
				bucket.head = current.Next
			} else {
				prev.Next = current.Next
			}
			h.size--
			return
		}
		prev = current
		current = current.Next
	}
}

// resize doubles the bucket list size and rehashes all keys
func (h *HashMap) resize() {
	newCapacity := h.capacity * 2
	newBuckets := NewBucketList(newCapacity)

	// Rehash all existing keys
	for _, bucket := range h.buckets {
		current := bucket.head
		for current != nil {
			newIndex := HashFunction(current.Key, newCapacity)
			newBucket := newBuckets[newIndex]

			// Move entry to new bucket
			newEntry := &BucketNode{Key: current.Key, Value: current.Value, Next: newBucket.head}
			newBucket.head = newEntry

			current = current.Next
		}
	}

	// Update HashMap metadata
	h.buckets = newBuckets
	h.capacity = newCapacity
}

func HashMapRun() {
	hm := NewHashMap(2)

	hm.Set("name", "Alice")
	hm.Set("age", 25)
	hm.Set("city", "New York")
	hm.Set("email", "alice@example.com")
	hm.Set("country", "USA")

	val, found := hm.Get("name")
	if found {
		fmt.Println("Name:", val)
	} else {
		fmt.Println("Name not found")
	}

}
