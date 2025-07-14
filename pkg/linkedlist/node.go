package linkedlist

type Node[T any] struct {
	Data T
	Next *Node[T]
}
