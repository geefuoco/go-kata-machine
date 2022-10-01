package katas


type LinkedList[T any] struct {
}

func NewLinkedList[T any]() *LinkedList[T] {
  return nil
}

func (l *LinkedList[T]) Append(val T) {
}

func (l *LinkedList[T]) Prepend(val T) {
}

func (l *LinkedList[T]) Get(idx int) T {
  var out T
  return out
}

func (l *LinkedList[T]) InsertAfter(idx int, val T) {
}

func (l *LinkedList[T]) RemoveAt(idx int) T {
  var out T
  return out
}

func (l *LinkedList[T]) Pop() T {
  var out T
  return out
}


func (l *LinkedList[T]) Head() *Node[T]{
  return nil
}

func (l *LinkedList[T]) Tail() *Node[T]{
  return nil
}

func (l *LinkedList[T]) Length() int {
  return -1
}




