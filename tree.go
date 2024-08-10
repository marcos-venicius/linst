package main

func Create[T any]() *Tree[T] {
	return &Tree[T]{
		root:    nil,
		current: nil,
	}
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
	} else if t.current.next != nil && t.current.prev != nil {
		t.addMiddle(node)
	} else if t.current.prev == nil {
		t.addBeginning(node)
	}

	if t.root.prev != nil {
		panic("invalid tree")
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

func (t *Tree[T]) Next() (*Node[T], error) {
	if t.current.next == nil {
		return nil, &NodeNotFound{
			msg: "next node not found",
		}
	}

	t.current = t.current.next

	return t.current, nil
}

func (t *Tree[T]) Node() *Node[T] {
	return t.current
}

func (t *Tree[T]) SelectRoot() {
	if t.root != nil {
		t.current = t.root
	}
}

func (t *Tree[T]) HasNext() bool {
	if t.current == nil {
		return false
	}

	return t.current.next != nil
}

func (t *Tree[T]) HasPrev() bool {
	if t.current == nil {
		return false
	}

	return t.current.prev != nil
}
