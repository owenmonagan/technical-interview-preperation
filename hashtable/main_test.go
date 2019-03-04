package hashtable

import (
	"fmt"
	"testing"

	"github.com/technical-interview-preperation/linkedlist"
)

func TestHash(t *testing.T) {
	key := "Zach McJohnson"
	tableSize := 1002
	h := hash(key, tableSize)
	if tableSize < h || h < 0 {
		t.Error("failed to get a hash within the range", h)
	}
	h2 := hash(key, tableSize)
	if h2 != h {
		t.Error("hash values not consistent", h, h2)
	}
}

func TestNew(t *testing.T) {
	tableSize := 1002
	table := New(tableSize)
	if len(table.hashSlice) != tableSize {
		t.Error("incorrect table size")
	}
	fmt.Println(table.hashSlice[0])
}

func TestInsert(t *testing.T) {
	tableSize := 10
	table := New(tableSize)
	key := "Zach McJohnson"
	h := hash(key, tableSize)
	node := linkedlist.Node{ID: key}
	insert(table, &node)
	if table.hashSlice[h] == (linkedlist.LinkedList{}) {
		t.Error("failed to insert")
	}
	if node, found := table.hashSlice[h].Search(key); found == false {
		t.Error("didn't find the node")
	} else {
		if node.GetID() != key {
			t.Error("didn't get the correct key")
		}
	}
}

func TestRemove(t *testing.T) {
	tableSize := 10
	table := New(tableSize)
	key := "Zach McJohnson"
	h := hash(key, tableSize)
	node := linkedlist.Node{ID: key}
	insert(table, &node)
	remove(table, key)
	if _, found := table.hashSlice[h].Search(key); found {
		t.Error("should not of found the node")
	}
}
