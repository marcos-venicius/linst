package main

func Create[T any]() *Tree[T] {
	return &Tree[T]{
		root:    nil,
		current: nil,
	}
}

func (t *Tree[T]) IsEmpty() bool {
	return t.root == nil
}

// adds a new node in the tree and return the pointer
func (t *Tree[T]) Add(data T) *Node[T] {
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

	return node
}

// Delete current node
func (t *Tree[T]) Delete() error {
	if t.root == nil {
		return &EmptyTreeError{
			msg: "Tree is empty",
		}
	}

	if t.current.next == nil && t.current.prev == nil {
		t.root = nil
		t.current = nil
		return nil
	} else if t.current.prev != nil && t.current.next != nil {
		t.current.next.prev, t.current.prev.next = t.current.prev, t.current.next
		t.current = t.current.next
		return nil
	} else if t.current.prev == nil && t.current.next != nil {
		t.current = t.current.next
		t.current.prev = nil
		t.root = t.current
		return nil
	} else if t.current.next == nil && t.current.prev != nil {
		t.current = t.current.prev
		t.current.next = nil
		t.root = t.current
		return nil
	}

	panic("Something went wrong")
}

// select and return prev node
func (t *Tree[T]) Prev() (*Node[T], error) {
	if t.current.prev == nil {
		return nil, &NodeNotFoundError{
			msg: "prev node not found",
		}
	}

	t.current = t.current.prev

	return t.current, nil
}

// select and return next node
func (t *Tree[T]) Next() (*Node[T], error) {
	if t.current.next == nil {
		return nil, &NodeNotFoundError{
			msg: "next node not found",
		}
	}

	t.current = t.current.next

	return t.current, nil
}

// returns current node
func (t *Tree[T]) Node() *Node[T] {
	return t.current
}

// select root node
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
