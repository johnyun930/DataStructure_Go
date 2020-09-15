package dataStructure

type LinkedStack struct {
	topNode    *Node
	numofEntry int
}

func (l *LinkedStack) Push(val Data) {
	newNode := &Node{val: val}
	newNode.next = l.topNode
	l.topNode = newNode
	l.numofEntry++

}

func (l *LinkedStack) Pop() *Node {
	removeNode := l.Peek()
	if removeNode == nil {
		return nil
	}
	l.topNode = l.topNode.next
	l.numofEntry--
	return removeNode
}

func (l *LinkedStack) Peek() *Node {
	if l.IsEmpty() {
		return nil
	} else {
		return l.topNode
	}
}

func (l *LinkedStack) IsEmpty() bool {
	if l.numofEntry == 0 {
		return true
	} else {
		return false
	}
}

func (l *LinkedStack) Clear() {
	l.topNode = nil
	l.numofEntry = 0
}
