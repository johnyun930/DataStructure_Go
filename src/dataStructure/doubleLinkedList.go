package dataStructure

type DoubleLinkedList struct {
	root       *Node
	tail       *Node
	numofEntry int
}

func (d *DoubleLinkedList) Add(val Data) bool {
	newNode := &Node{val: val}

	if d.root == nil {
		d.root = newNode
		d.tail = newNode
	} else {
		d.tail.next = newNode
		newNode.previous = d.tail
		d.tail = newNode
	}
	d.numofEntry++
	return true

}

func (d *DoubleLinkedList) AddAt(index int, val Data) bool {
	newNode := &Node{val: val}
	if index <= d.numofEntry && index > 0 {
		if index == 0 {
			if d.Size() == 0 {
				d.root = newNode
				d.tail = newNode
			} else {
				newNode.next = d.root
				d.root.previous = newNode
				d.root = newNode
			}
		} else if index == d.numofEntry {
			d.tail.next = newNode
			newNode.previous = d.tail
			d.tail = newNode
		} else {
			beforeNode := d.GetNodeAt(index - 1)
			nextNode := beforeNode.next
			beforeNode.next = newNode
			newNode.previous = beforeNode
			newNode.next = nextNode
			nextNode.previous = newNode
		}
		d.numofEntry++
		return true
	} else {
		return false
	}
}

func (d *DoubleLinkedList) RemoveNodeAt(entry int) *Node {
	var removeNode *Node = nil
	if d.IsInRange(entry) {
		if entry == 0 {
			removeNode = d.root
			d.root = d.root.next
			d.root.previous = nil
		} else if entry == d.numofEntry-1 {
			removeNode = d.tail
			d.tail = d.tail.previous
			d.tail.next = nil
		} else {
			removeNode = d.GetNodeAt(entry)
			beforeNode := removeNode.previous
			nextNode := removeNode.next
			beforeNode.next = nextNode
			nextNode.previous = beforeNode
			removeNode.previous = nil
			removeNode.next = nil

			d.numofEntry--

		}
	}
	return removeNode
}

func (d *DoubleLinkedList) RemoveNode(val Data) bool {
	index := d.FindNode(val)
	if index < 0 {
		return false
	} else {
		if d.RemoveNodeAt(index) != nil {
			return true
		} else {
			return false
		}
	}
}

func (d *DoubleLinkedList) FindNode(val Data) int {
	var index int = -1
	currentNode := d.root
	for i := 0; i < d.Size(); i++ {
		if currentNode.val == val {
			index = i
			break
		}
	}
	return index
}

func (d *DoubleLinkedList) GetNodeAt(index int) *Node {
	if d.IsInRange(index) {
		return nil
	}
	if index == 0 {
		return d.root
	}
	if index == d.numofEntry-1 {
		return d.tail
	}
	node := d.root
	for i := 0; i < index; i++ {
		node = node.next
	}
	return node

}

func (d *DoubleLinkedList) Clear() {
	d.root = nil
	d.tail = nil
	d.numofEntry = 0

}

func (d *DoubleLinkedList) Size() int {
	return d.numofEntry
}

func (d *DoubleLinkedList) IsInRange(index int) bool {
	if index >= d.numofEntry || index < 0 {
		return false
	} else {
		return true
	}
}
