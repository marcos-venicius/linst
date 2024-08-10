package main

type Node[T any] struct {
	data T
	prev *Node[T]
	next *Node[T]
}

type Tree[T any] struct {
	root    *Node[T]
	current *Node[T]
}

func CreateTree[T any]() *Tree[T] {
	return &Tree[T]{
		root:    nil,
		current: nil,
	}
}

func (t *Tree[T]) addFirst(node *Node[T]) {
	t.root = node
	t.current = node
}

func (t *Tree[T]) addLast(node *Node[T]) {
	node.prev = t.current
	t.current.next = node

	t.current = node
}

func (t *Tree[T]) Add(data T) {
	node := &Node[T]{
		data: data,
		prev: nil,
		next: nil,
	}

	if t.root == nil {
		t.addFirst(node)
	} else if t.current.next == nil {
		t.addLast(node)
	}
}

func (t *Tree[T]) Prev() (*Node[T], error) {
	if t.current.prev == nil {
		return nil, &NodeNotFound{
			msg: "prev node not found",
		}
	}

	t.current = t.current.prev

	return t.current, nil
}
