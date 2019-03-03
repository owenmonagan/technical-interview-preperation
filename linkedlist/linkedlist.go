package linkedlist

type LinkedList struct {
	start *Node
}

type Node struct {
	id    string
	next  *Node
	phone int
}

func (list *LinkedList) insert(new *Node) {
	if *list == (LinkedList{}) {
		list.start = new
	} else {
		current := list.start
		inserted := false
		for !inserted {
			if current.id == new.id {
				list.start.phone = new.phone
				inserted = true
			} else if current.next == nil {
				current.next = new
				inserted = true
			} else {
				current = current.next
			}
		}
	}
	return
}

func (list *LinkedList) remove(id string) bool {
	if *list == (LinkedList{}) {
		return false
	}
	if list.start.id == id {
		list.start = list.start.next
		return true
	}
	current := list.start
	for {
		if current.next == nil {
			return false
		} else if current.next.id == id {
			current.next = current.next.next
			return true
		} else {
			current = current.next
		}
	}

}
