package main

type NodeNotFoundError struct {
	msg string
}

func (e *NodeNotFoundError) Error() string {
	return e.msg
}

type EmptyTreeError struct {
	msg string
}

func (e *EmptyTreeError) Error() string {
	return e.msg
}
