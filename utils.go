package linst

func (t *Tree[T]) addBeginning(node *Node[T]) {
	t.current.prev = node
	node.next = t.current
	t.current = node
	t.root = node
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

func (t *Tree[T]) addMiddle(node *Node[T]) {
	node.prev = t.current.prev
	t.current.prev.next = node
	node.next = t.current
	t.current.prev = node
}
