package katas


type LinkedList[T any] struct {
}

func NewLinkedList[T any]() *LinkedList[T] {
}

func (l *LinkedList[T]) Append(val T) {
}

func (l *LinkedList[T]) Prepend(val T) {
}

func (l *LinkedList[T]) Get(idx uint) T {
}

func (l *LinkedList[T]) InsertAfter(idx uint, val T) {
}

func (l *LinkedList[T]) RemoveAt(idx uint) T {
}

func (l *LinkedList[T]) Pop() T {
}


func (l *LinkedList[T]) Head() *Node[T]{
}

func (l *LinkedList[T]) Tail() *Node[T]{
}

func (l *LinkedList[T]) Length() uint {
}




