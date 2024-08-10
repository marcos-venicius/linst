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
