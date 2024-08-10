package main

type NodeNotFound struct {
	msg string
}

func (e *NodeNotFound) Error() string {
	return e.msg
}
