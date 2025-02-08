package linkedlist

import "errors"

type (
	Node struct {
		Value interface{}
		next  *Node
		prev  *Node
	}

	List struct {
		head *Node
		tail *Node
	}
)

func nodeOf(val interface{}) *Node {
	return &Node{Value: val}
}

func listOf(head, tail *Node) *List {
	return &List{head: head, tail: tail}
}

func NewList(elements ...interface{}) *List {
	var head *Node
	var tail *Node
	var current *Node
	for i, v := range elements {
		node := nodeOf(v)
		if i == 0 {
			head = node
		}
		if i == len(elements)-1 {
			tail = node
		}
		if current != nil {
			current.next = node
			node.prev = current
		}
		current = node
	}
	return listOf(head, tail)
}

func (n *Node) Next() *Node {
	return n.next
}

func (n *Node) Prev() *Node {
	return n.prev
}

func (l *List) Unshift(v interface{}) {
	node := nodeOf(v)
	prevHead := l.head
	if prevHead == nil {
		l.head = node
		l.tail = node
		return
	}

	prevHead.prev = node
	node.next = prevHead
	l.head = node
}

func (l *List) Push(v interface{}) {
	node := nodeOf(v)
	prevTail := l.tail
	if prevTail == nil {
		l.tail = node
	}
	if l.head == nil {
		l.head = node
		return
	}

	prevTail.next = node
	node.prev = prevTail
	l.tail = node
}

func (l *List) Shift() (interface{}, error) {
	if l.head == nil {
		return nil, errors.New("no head")
	}
	n := l.head
	if n == l.tail {
		l.head = nil
		l.tail = nil
	} else {
		l.head = n.next
		l.head.prev = nil
	}
	return n.Value, nil
}

func (l *List) Pop() (interface{}, error) {
	if l.tail == nil {
		return nil, errors.New("no tail")
	}
	n := l.tail
	if n == l.head {
		l.head = nil
		l.tail = nil
	} else {
		l.tail = n.prev
		l.tail.next = nil
	}
	return n.Value, nil
}

func (l *List) Reverse() {
	nodes := []*Node{}

	current := l.tail
	for current != nil {
		nodes = append(nodes, current)
		current = current.prev
	}

	current = nil
	for i, v := range nodes {
		v.next = nil
		v.prev = nil
		if i == 0 {
			l.head = v
		}
		if i == len(nodes)-1 {
			l.tail = v
		}
		if current != nil {
			current.next = v
			v.prev = current
		}
		current = v
	}
}

func (l *List) First() *Node {
	return l.head
}

func (l *List) Last() *Node {
	return l.tail
}
