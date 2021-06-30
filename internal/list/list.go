package list

type Node struct {
	buf  []byte
	next *Node
}

type List struct {
	head *Node
	tail *Node
}

func (l *List) Append(buf []byte) {
	node := &Node{
		buf: buf,
	}
	if l.head == nil {
		l.head = node
	}
	if l.tail == nil {
		l.tail = l.head
	}
	l.tail.next = node
	l.tail = node
}

func (l *List) Shrink() {
	if l.head == nil {
		return
	}
	if l.head == l.tail {
		l.head = nil
		l.tail = nil
		return
	}
	l.head = l.head.next
}
