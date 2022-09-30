package katas

type DoublyLinkedList[T any] struct {
}

func NewDoublyLinkedList[T any]() *DoublyLinkedList[T] {
}

func (d *DoublyLinkedList[T]) Append(val T) {
}

func (d *DoublyLinkedList[T]) Prepend(val T) {
}

func (d *DoublyLinkedList[T]) InsertAt(idx int, val T) {
}

func (d *DoublyLinkedList[T]) RemoveAt(idx int) T{
}

func (d *DoublyLinkedList[T]) Deque(idx int) T {
}

func (d *DoublyLinkedList[T]) Pop() T {
}

func (d *DoublyLinkedList[T]) Get(idx int) T {
}


func (d *DoublyLinkedList[T]) Length() int {
}

func (d *DoublyLinkedList[T]) Head() *Node[T] {
}
func (d *DoublyLinkedList[T]) Tail() *Node[T] {
}
