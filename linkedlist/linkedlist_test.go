package linkedlist

import (
	"testing"
)

func TestNodeInsert(t *testing.T) {
	list := LinkedList{}

	id := "John Smith"
	phone := 7322983360
	node := &Node{id: id, phone: phone}
	list.insert(node)
	if list.start.id != id {
		t.Error("didn't get the correct id")
	}
	if list.start.phone != phone {
		t.Error("didn't get the correct phone number")
	}
	phone = 92321355
	node = &Node{id: id, phone: phone}
	list.insert(node)
	if node.phone != phone {
		t.Error("didn't get the correct phone number")
	}
	phone = 9232135523
	id = "Swag Micheal Junior"

	node = &Node{id: id, phone: phone}
	list.insert(node)
	if list.start.next.phone != phone {
		t.Error("didn't get the correct phone number")
	}
	node = &Node{id: "Zach  Paul", phone: 8123123}
	list.insert(node)
	node = &Node{id: "Jake Monagan", phone: 123123123}
	list.insert(node)
	if list.start.next.next.next == nil {
		t.Error("should not of gotten nil node")
	}
	if list.start.next.next.next.next != nil {
		t.Error("should of gotten nil node")
	}
}

func TestRemove(t *testing.T) {
	list := LinkedList{}
	removeID := "Dave Donnelly"
	replacementNodeID := "Tim Burke"
	list.insert(&Node{id: removeID, phone: 2319203813})
	if !list.remove(removeID) {
		t.Error("should of removed ", removeID)
	}
	list.insert(&Node{id: "Zach  Paul", phone: 8123123})
	list.insert(&Node{id: "Jake Monagan", phone: 123123123})
	list.insert(&Node{id: "Tom Pitt", phone: 345235})
	list.insert(&Node{id: removeID, phone: 2319203813})
	list.insert(&Node{id: replacementNodeID, phone: 345235})
	if !list.remove(removeID) {
		t.Error("should of removed ", removeID)
	}
	if list.start.next.next.next == nil || list.start.next.next.next.id != replacementNodeID {
		t.Error("should of had time burke as the element")
	}
	if list.remove("random string") {
		t.Error("should not of removed ", removeID)
	}
}
