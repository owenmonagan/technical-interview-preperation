package linkedlist

type LinkedList struct {
	start *Node
}

type Node struct {
	ID    string
	next  *Node
	phone int
}

func (node Node) GetPhone() int {
	return node.phone
}

func (node Node) GetID() string {
	return node.ID
}

func (list *LinkedList) Search(id string) (Node, bool) {
	if *list == (LinkedList{}) {
		return Node{}, false
	}
	current := list.start
	for {
		if current == nil {
			return Node{}, false
		}
		if current.ID == id {
			return *current, true
		}
	}
}

func (list *LinkedList) Insert(new *Node) {
	if *list == (LinkedList{}) {
		list.start = new
	} else {
		current := list.start
		inserted := false
		for !inserted {
			if current.ID == new.ID {
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

func (list *LinkedList) Remove(ID string) bool {
	if *list == (LinkedList{}) {
		return false
	}
	if list.start.ID == ID {
		list.start = list.start.next
		return true
	}
	current := list.start
	for {
		if current.next == nil {
			return false
		} else if current.next.ID == ID {
			current.next = current.next.next
			return true
		} else {
			current = current.next
		}
	}

}
