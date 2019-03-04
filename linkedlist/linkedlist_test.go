package linkedlist

import (
	"testing"
)

func TestNodeInsert(t *testing.T) {
	list := LinkedList{}

	ID := "John Smith"
	phone := 7322983360
	node := &Node{ID: ID, phone: phone}
	list.Insert(node)
	if list.start.ID != ID {
		t.Error("dIDn't get the correct ID")
	}
	if list.start.phone != phone {
		t.Error("dIDn't get the correct phone number")
	}
	phone = 92321355
	node = &Node{ID: ID, phone: phone}
	list.Insert(node)
	if node.phone != phone {
		t.Error("dIDn't get the correct phone number")
	}
	phone = 9232135523
	ID = "Swag Micheal Junior"

	node = &Node{ID: ID, phone: phone}
	list.Insert(node)
	if list.start.next.phone != phone {
		t.Error("dIDn't get the correct phone number")
	}
	node = &Node{ID: "Zach  Paul", phone: 8123123}
	list.Insert(node)
	node = &Node{ID: "Jake Monagan", phone: 123123123}
	list.Insert(node)
	if list.start.next.next.next == nil {
		t.Error("should not of gotten nil node")
	}
	if list.start.next.next.next.next != nil {
		t.Error("should of gotten nil node")
	}
}

func TestRemove(t *testing.T) {
	list := LinkedList{}
	RemoveID := "Dave Donnelly"
	replacementNodeID := "Tim Burke"
	list.Insert(&Node{ID: RemoveID, phone: 2319203813})
	if !list.Remove(RemoveID) {
		t.Error("should of Removed ", RemoveID)
	}
	list.Insert(&Node{ID: "Zach  Paul", phone: 8123123})
	list.Insert(&Node{ID: "Jake Monagan", phone: 123123123})
	list.Insert(&Node{ID: "Tom Pitt", phone: 345235})
	list.Insert(&Node{ID: RemoveID, phone: 2319203813})
	list.Insert(&Node{ID: replacementNodeID, phone: 345235})
	if !list.Remove(RemoveID) {
		t.Error("should of Removed ", RemoveID)
	}
	if list.start.next.next.next == nil || list.start.next.next.next.ID != replacementNodeID {
		t.Error("should of had time burke as the element")
	}
	if list.Remove("random string") {
		t.Error("should not of Removed ", RemoveID)
	}
}
