package hashtable

import (
	"hash/fnv"
	"sync"

	"github.com/technical-interview-preperation/linkedlist"
)

type Hashtable struct {
	hashSlice []linkedlist.LinkedList
	lock      sync.RWMutex
}

func New(tableSize int) *Hashtable {
	table := Hashtable{}
	table.hashSlice = make([]linkedlist.LinkedList, tableSize)
	return &table
}

func hash(key string, tableSize int) int {
	h := 0
	keyHash := fnv.New32()
	keyHash.Write([]byte(key))
	for i := 0; i < len(key); i++ {
		h = (31*h + int(keyHash.Sum32())) % tableSize
	}
	return h
}

func insert(table *Hashtable, node *linkedlist.Node) {
	table.lock.Lock()
	defer table.lock.Unlock()
	h := hash(node.ID, len(table.hashSlice))
	table.hashSlice[h].Insert(node)
	return
}

func remove(table *Hashtable, key string) bool {
	table.lock.Lock()
	defer table.lock.Unlock()
	h := hash(key, len(table.hashSlice))
	return table.hashSlice[h].Remove(key)
}
