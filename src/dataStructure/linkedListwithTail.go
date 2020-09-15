package dataStructure

type LinkedListwithTail struct {
	root       *Node
	tail       *Node
	numofEntry int
}

func (l *LinkedListwithTail) Add(val Data) bool {
	newNode := &Node{val: val}
	if l.root == nil {
		l.root = newNode
		l.tail = l.root
	} else {
		l.tail.next = newNode
		l.tail = l.tail.next
	}
	l.numofEntry++
	return true
}
func (l *LinkedListwithTail) AddAt(entry int, val Data) bool {
	newNode := &Node{val: val}

	if l.numofEntry >= entry && entry >= 0 {
		if entry == 0 {
			if l.Size() == 0 {
				l.tail = newNode
			} else {
				newNode.next = l.root
			}
			l.root = newNode
			l.numofEntry++
			return true
		} else if entry == l.numofEntry {
			beforeNode := l.GetNodeAt(entry - 1)
			beforeNode.next = newNode
			l.tail = newNode
		} else {
			beforeNode := l.GetNodeAt(entry - 1)
			nextnode := l.GetNodeAt(entry)
			beforeNode.next = newNode.next
			newNode.next = nextnode
		}
		l.numofEntry++
		return true
	} else {
		return false
	}

}

func (l *LinkedListwithTail) FindNode(val Data) int {
	var index int = -1
	node := l.root
	for i := 0; i < l.numofEntry; i++ {
		if node.val == val {
			index = i
			break
		}
		node = node.next

	}
	return index

}

func (l *LinkedListwithTail) GetNodeAt(index int) *Node {
	if l.IsInRange(index) {
		return nil
	}
	if index == 0 {
		return l.root
	}
	if index == l.numofEntry-1 {
		return l.tail
	}
	node := l.root
	for i := 0; i < index; i++ {
		node = node.next
	}
	return node

}
func (l *LinkedListwithTail) RemoveNodeAt(entry int) *Node {
	currentNode := l.GetNodeAt(entry)
	if currentNode != nil {
		if entry == 0 {
			l.root = nil
			if l.numofEntry == 1 {
				l.tail = nil
			}
			l.numofEntry--
			return currentNode
		}
		beforeNode := l.GetNodeAt(entry - 1)
		if entry == l.numofEntry-1 {
			beforeNode.next = nil
			l.tail = beforeNode
		} else {
			beforeNode.next = currentNode.next
		}
		l.numofEntry--
		return currentNode
	} else {
		return nil
	}
}

func (l *LinkedListwithTail) RemoveNode(val Data) bool {
	index := l.FindNode(val)
	if index == -1 {
		return false
	}

	if l.RemoveNodeAt(index) != nil {
		return true
	} else {
		return false
	}

}

func (l *LinkedListwithTail) Clear() {
	l.root = nil
	l.tail = nil
	l.numofEntry = 0
}

func (l *LinkedListwithTail) IsInRange(entry int) bool {
	if entry >= l.numofEntry || entry < 0 {
		return false
	} else {
		return true
	}
}

func (l *LinkedListwithTail) Size() int {
	return l.numofEntry
}

func (l *LinkedListwithTail) ToArray() []Data {
	var arr []Data
	currentNode := l.root

	for i := 0; currentNode != l.tail; i++ {
		arr[i] = currentNode.val
		currentNode = currentNode.next
	}
	return arr
}
