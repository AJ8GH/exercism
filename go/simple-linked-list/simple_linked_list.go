package linkedlist

import "errors"

type (
	List struct {
		head *Node
		size int
	}

	Node struct {
		element int
		next    *Node
	}
)

func New(elements []int) *List {
	l := &List{size: len(elements)}
	if len(elements) == 0 {
		return l
	}
	head := &Node{element: elements[0]}
	l.head = head
	current := head
	for i := 1; i < len(elements); i++ {
		current.next = &Node{element: elements[i]}
		current = current.next
	}
	return l
}

func (l *List) Size() int {
	return l.size
}

func (l *List) Push(element int) {
	if l.size == 0 {
		l.head = &Node{element: element}
	} else {
		current := l.head
		for current.next != nil {
			current = current.next
		}
		current.next = &Node{element: element}
	}
	l.size++
}

func (l *List) Pop() (int, error) {
	if l.size == 0 {
		return 0, errors.New("no stuff")
	}
	if l.size == 1 {
		out := l.head
		l.head = nil
		l.size--
		return out.element, nil
	}
	var prev *Node
	current := l.head
	for current.next != nil {
		prev = current
		current = current.next
	}
	out := current.element
	prev.next = nil

	l.size--
	return out, nil
}

func (l *List) Array() (out []int) {
	current := l.head
	for current != nil {
		out = append(out, current.element)
		current = current.next
	}
	return out
}

func (l *List) Reverse() *List {
	if l.size < 2 {
		return l
	}

	nodes := []*Node{}
	current := l.head
	for current != nil {
		nodes = append(nodes, current)
		current = current.next
	}

	head := nodes[len(nodes)-1]
	head.next = nil
	current = head
	for i := len(nodes) - 2; i >= 0; i-- {
		n := nodes[i]
		n.next = nil
		current.next = n
		current = n
	}
	l.head = head
	return l
}
