package util


type LinkedList[T any] struct {
  head *Node[T]
  tail *Node[T]
  length uint
}

func NewLinkedList[T any]() *LinkedList[T] {
  return &LinkedList[T]{nil, nil, 0}
}

func (l *LinkedList[T]) Append(n *Node[T]) {
  if l.head == nil {
    l.head = n
    l.tail = l.head
  } else {
    l.tail.Next = n
    l.tail = l.tail.Next
  }
  l.length++
}

func (l *LinkedList[T]) Prepend(n *Node[T]) {
  if l.head != nil {
    temp := l.head
    l.head = n;
    l.head.Next = temp
  } else {
    l.head = n
    l.tail = l.head
  }
  l.length++
}

func (l *LinkedList[T]) Get(idx uint) *Node[T] {
  return l.get(idx)
}

//Singly linked list can only insert after
func (l *LinkedList[T]) InsertAfter(idx uint, node *Node[T]) {
  old_node := l.get(idx)
  temp := old_node.Next
  old_node.Next = node
  node.Next = temp
  l.length++
}

func (l *LinkedList[T]) RemoveAt(idx uint) *Node[T] {
  if idx == 0 {
    out := l.head
    l.head = l.head.Next
    l.length--
    if l.length == 0 {
      l.head = nil
      l.tail = nil
    }
    return out
  }
  if idx == l.length-1 {
    out := l.tail
    prev := l.get(idx-1)
    prev.Next = nil
    l.tail = prev
    l.length--
    return out 
  }
  prev := l.get(idx-1)
  out := prev.Next
  prev.Next = out.Next
  l.length--
  return out 
}

func (l *LinkedList[T]) Pop() *Node[T] {
  return l.RemoveAt(l.length-1)
}

func (l *LinkedList[T]) get(idx uint) *Node[T] {
  if idx >= l.length {
    panic("Index outside of bounds")
  }
  var num uint
  current := l.head
  for current.Next != nil && num < idx {
    current = current.Next
    num++
  }
  return current
}

func (l *LinkedList[T]) Head() *Node[T]{
  return l.head
}

func (l *LinkedList[T]) Tail() *Node[T]{
  return l.tail
}

func (l *LinkedList[T]) Length() uint {
  return l.length
}




