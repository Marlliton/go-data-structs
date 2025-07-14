package linkedlist

import (
	"fmt"
	"strings"
)

type LinkedList[T any] struct {
	head  *Node[T]
	count int
}

func (ll *LinkedList[T]) Prepend(data T) {
	n := &Node[T]{Data: data, Next: ll.head}
	ll.head = n
	ll.count++
}

func (ll *LinkedList[T]) Append(data T) {
	n := &Node[T]{Data: data, Next: nil}
	if ll.head == nil {
		ll.head = n
		ll.count++
		return
	}

	current := ll.head
	for current.Next != nil {
		current = current.Next
	}

	current.Next = n
	ll.count++
}

func (ll *LinkedList[T]) GetAt(index int) (T, bool) {
	n, ok := ll.getNodeAt(index)
	if !ok {
		var zero T
		return zero, ok
	}
	return n.Data, ok
}

func (ll *LinkedList[T]) Find(fn func(T) bool) (T, bool) {
	var data T
	current := ll.head
	for current != nil {
		ok := fn(current.Data)
		if ok {
			data = current.Data
			return data, true
		}
		current = current.Next
	}

	return data, false
}

func (ll *LinkedList[T]) IndexOf(fn func(T) bool) int {
	current := ll.head
	count := 0
	for current != nil {
		if fn(current.Data) {
			return count
		}
		count++
		current = current.Next
	}

	return -1
}

func (ll *LinkedList[T]) InsertAt(data T, index int) bool {
	if index == 0 {
		ll.Prepend(data)
		return true
	}

	if index >= ll.count {
		ll.Append(data)
		return true
	}

	previous, ok := ll.getNodeAt(index - 1)
	if !ok {
		return false
	}

	n := &Node[T]{Data: data, Next: previous.Next}
	previous.Next = n
	ll.count++
	return true
}

func (ll *LinkedList[T]) RemoveAt(index int) (T, bool) {
	var zero T
	if index >= 0 && index < ll.count {
		node := ll.head
		if index == 0 {
			ll.head = node.Next
			ll.count--
			return node.Data, true
		}

		previous, ok := ll.getNodeAt(index - 1) // elemento anterior ao index
		if !ok {
			return zero, false
		}
		node = previous.Next // elemento exato ao index (cujo queremos remover)
		// var previous *Node[T]
		// for i := range index {
		// 	fmt.Println("removendo o indext", index, "item atual", i)
		// 	previous = current     // o previous é o nó de index-1
		// 	current = current.next // quanto chaga ao index o current é o nó de indice index
		// }

		previous.Next = node.Next // previous.next é o elemento que quremos remver. current.next é o elemento index+1. Dessa forma o elemento == index fica perdido na memória.
		ll.count--
		return node.Data, true
	}

	return zero, false
}

func (ll *LinkedList[T]) IsEmpyt() bool {
	return ll.head == nil
}

func (ll *LinkedList[T]) Len() int {
	return ll.count
}

func (ll *LinkedList[T]) getNodeAt(index int) (*Node[T], bool) {
	var zero *Node[T]
	if index < 0 || index > ll.count-1 {
		return zero, false
	}
	if index == 0 {
		return ll.head, true
	}

	current := ll.head
	for range index {
		current = current.Next
	}

	return current, true
}

func (ll *LinkedList[T]) String() string {
	if ll.head == nil {
		return "[]"
	}

	var st strings.Builder
	st.WriteString("[")
	current := ll.head
	for current != nil {
		st.WriteString(fmt.Sprintf("%v", current.Data))
		if current.Next != nil {
			st.WriteString(", ")
		}
		current = current.Next
	}
	st.WriteString("]")

	return st.String()
}
