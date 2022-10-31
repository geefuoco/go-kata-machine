package katas

type DoublyLinkedList[T any] struct {
}

func NewDoublyLinkedList[T any]() *DoublyLinkedList[T] {
  return nil
}

func (d *DoublyLinkedList[T]) Append(val T) {
}

func (d *DoublyLinkedList[T]) Prepend(val T) {
}

func (d *DoublyLinkedList[T]) InsertAt(idx int, val T) {
}

func (d *DoublyLinkedList[T]) RemoveAt(idx int) T{
  var out T
  return out
}

func (d *DoublyLinkedList[T]) Deque() T {
  var out T
  return out
}

func (d *DoublyLinkedList[T]) Pop() T {
  var out T
  return out
}

func (d *DoublyLinkedList[T]) Get(idx int) T {
  var out T
  return out
}


func (d *DoublyLinkedList[T]) Length() int {
  return -1
}

func (d *DoublyLinkedList[T]) Head() *Node[T] {
  return nil
}
func (d *DoublyLinkedList[T]) Tail() *Node[T] {
  return nil
}
