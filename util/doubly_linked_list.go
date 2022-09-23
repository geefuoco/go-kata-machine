package util

type DoublyLinkedList[T any] struct {
  head *Node[T]
  tail *Node[T]
  length int
}

func NewDoublyLinkedList[T any]() *DoublyLinkedList[T] {
  return &DoublyLinkedList[T]{}
}

func (d *DoublyLinkedList[T]) Append(val T) {
  node := &Node[T]{Val: val}
  d.length++
  if d.head == nil {
    d.head = node
    d.tail = node
  } else {
    d.tail.Next = node
    node.Prev = d.tail
    d.tail = node
  }
}

func (d *DoublyLinkedList[T]) Prepend(val T) {
  node := &Node[T]{Val: val}
  d.length++
  if d.head == nil {
    d.head = node
    d.tail = node
  } else {
    d.head.Prev=  node
    d.head = node
  }
}

func (d *DoublyLinkedList[T]) InsertAt(idx int, val T) {
  if idx == 0 {
    d.Prepend(val)
    return
  }
  node := d.get(idx)
  newNode := &Node[T]{Val: val}
  if node.Prev != nil {
    node.Prev.Next = newNode
    newNode.Prev = node.Prev
  }
  node.Prev = newNode
  newNode.Next = node
  d.length++
}

func (d *DoublyLinkedList[T]) RemoveAt(idx int) T{
  node := d.get(idx)
  if idx != 0 {
    node.Prev.Next = node.Next
  }
  if node.Next != nil {
    node.Next.Prev = node.Prev
  }
  if d.length == 1 {
    d.length = 0
    d.head = nil
    d.tail = nil
  } else {
    d.length--
  }
  return node.Val
}

func (d *DoublyLinkedList[T]) Deque(idx int) T {
  return d.RemoveAt(0)
}

func (d *DoublyLinkedList[T]) Pop() T {
  return d.RemoveAt(d.length-1)
}

func (d *DoublyLinkedList[T]) Get(idx int) T {
  var val T
  node := d.get(idx)
  if node != nil{
    val = node.Val
  }
  return val
}

func (d *DoublyLinkedList[T]) get(idx int) *Node[T]{
  if idx >= d.length {
    panic("Out of bounds")
  }
  front := idx 
  back := d.length-1 - idx
  var current *Node[T]
  var count int
  if front < back {
    count = 0
    current = d.head
    for current.Next != nil && count < idx {
      current = current.Next
      count++
    }
  } else {
    count = d.length-1
    current = d.tail
    for current.Prev != nil && count > idx {
      current = current.Prev
      count--
    }
  }
  return current
}

func (d *DoublyLinkedList[T]) Length() int {
  return d.length
}

func (d *DoublyLinkedList[T]) Head() *Node[T] {
  return d.head
}
func (d *DoublyLinkedList[T]) Tail() *Node[T] {
  return d.tail
}
