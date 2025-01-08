package linkedlist

import "errors"

// Define List and Node types here.
type List struct {
	Head *Node
	Tail *Node
}

type Node struct {
	Value interface{}
	N     *Node
	P     *Node
}

// Note: The tests expect Node type to include an exported field with name Value to pass.

func NewList(elements ...interface{}) *List {
	list := &List{}
	for _, element := range elements {
		list.Push(element)
	}
	return list
}

func (n *Node) Next() *Node {
	return n.N
}

func (n *Node) Prev() *Node {
	return n.P
}

func (l *List) Unshift(v interface{}) {
	node := &Node{
		Value: v,
	}
	if l.Head == nil {
		l.Head = node
		l.Tail = node
	} else {
		node.N = l.Head
		node.N.P = node
		l.Head = node
	}
}

func (l *List) Push(v interface{}) {
	node := &Node{
		Value: v,
	}
	if l.Head == nil {
		l.Head = node
	}
	if l.Tail == nil {
		l.Tail = node
	} else {
		l.Tail.N = node
		node.P = l.Tail
		l.Tail = node
	}
}

func (l *List) Shift() (interface{}, error) {
	if l.Head == nil {
		return nil, errors.New("list is empty")
	}
	node := l.Head
	if node.N == nil {
		l.Tail = nil
		l.Head = nil
	} else {
		node.N.P = nil
		l.Head = node.N
	}
	return node.Value, nil
}

func (l *List) Pop() (interface{}, error) {
	if l.Tail == nil || l.Head == nil {
		return nil, errors.New("List is empty")
	}
	node := l.Tail
	l.Tail = node.P
	if l.Tail == nil {
		l.Head = nil
	} else {
		l.Tail.N = nil
	}
	return node.Value, nil
}

func (l *List) Reverse() {
	var prev *Node
	curr := l.Head
	l.Head, l.Tail = l.Tail, l.Head
	for curr != nil {
		prev = curr.P
		curr.P = curr.N
		curr.N = prev
		curr = curr.P
	}
}

func (l *List) First() *Node {
	return l.Head
}

func (l *List) Last() *Node {
	return l.Tail
}
