package dataStructure

type LinkedList struct {
	root       *Node
	numofEntry int
}

func MakeLinkedList() *LinkedList {
	return &LinkedList{root: nil, numofEntry: 0}
}

func (l *LinkedList) Add(val Data) bool {
	newNode := &Node{val: val}
	if l.root == nil {
		l.root = newNode
	} else {
		var endNode *Node = l.root
		for i := 1; i < l.numofEntry; i++ {
			endNode = endNode.next
		}
		endNode.next = newNode
	}
	l.numofEntry++
	return true
}

func (l *LinkedList) AddAt(entry int, val Data) bool {
	newNode := &Node{val: val}

	if l.numofEntry >= entry && entry >= 0 {
		if entry == 0 {
			newNode.next = l.root
			l.root = newNode.next
		} else if entry == l.numofEntry {
			beforeNode := l.GetNodeAt(entry - 1)
			beforeNode.next = newNode
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

func (l *LinkedList) RemoveNodeAt(entry int) *Node {
	currentNode := l.GetNodeAt(entry)
	if currentNode != nil {
		if entry == 0 {
			l.root = nil
			l.numofEntry--
			return currentNode
		}
		beforeNode := l.GetNodeAt(entry - 1)
		if entry == l.numofEntry-1 {
			beforeNode.next = nil
		} else {
			beforeNode.next = currentNode.next
		}
		l.numofEntry--
		return currentNode
	} else {
		return nil
	}

}

func (l *LinkedList) RemoveNode(val Data) bool {
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

func (l *LinkedList) FindNode(val Data) int {
	var entry int = -1
	node := l.root
	for i := 0; i < l.numofEntry; i++ {
		if node.val == val {
			entry = i
			break
		}
		node = node.next

	}
	return entry

}

func (l *LinkedList) GetNodeAt(index int) *Node {
	if l.IsInRange(index) {
		return nil
	}
	node := l.root
	for i := 0; i < index; i++ {
		node = node.next
	}
	return node
}

func (l *LinkedList) Replace(index int, data Data) Data {
	node := l.GetNodeAt(index)
	beforeData := node.val
	node.val = data
	return beforeData
}

func (l *LinkedList) IsInRange(index int) bool {
	if index >= l.numofEntry || index < 0 {
		return false
	} else {
		return true
	}
}

func (l *LinkedList) Clear() {
	l.root = nil
	l.numofEntry = 0
}

func (l *LinkedList) Size() int {
	return l.numofEntry
}

func (l *LinkedList) ToArray() []Data {
	var arr []Data

	currentNode := l.root

	for i := 0; i < l.Size(); i++ {
		arr[i] = currentNode.val
		currentNode = currentNode.next
	}
	return arr
}

func (l *LinkedList) PrintData() {
	currentNode := l.root
	for i := 0; i < l.Size(); i++ {
		println(currentNode.val.Val)
		currentNode = currentNode.next
	}
}
