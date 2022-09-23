package util


type LinkedList[T any] struct {
  head *Node[T]
  tail *Node[T]
  length uint
}

func (l *LinkedList[T]) Append(val T) {
  l.append(&Node[T]{Val: val, Next: nil, Prev: nil})
}

func (l *LinkedList[T]) Prepend(val T) {
  l.prepend(&Node[T]{Val: val, Next: nil, Prev: nil})
}

func (l *LinkedList[T]) Get(idx uint) T {
  return l.get(idx).Val
}

func (l *LinkedList[T]) InsertAfter(idx uint, val T) {
  l.insertAfter(idx, &Node[T]{Val: val})
}

func (l *LinkedList[T]) RemoveAt(idx uint) T {
  return l.removeAt(idx).Val
}

func (l *LinkedList[T]) Pop() T {
  return l.removeAt(l.length-1).Val
}

func NewLinkedList[T any]() *LinkedList[T] {
  return &LinkedList[T]{nil, nil, 0}
}

func (l *LinkedList[T]) append(n *Node[T]) {
  if l.head == nil {
    l.head = n
    l.tail = l.head
  } else {
    l.tail.Next = n
    l.tail = l.tail.Next
  }
  l.length++
}

func (l *LinkedList[T]) prepend(n *Node[T]) {
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

//Singly linked list can only insert after
func (l *LinkedList[T]) insertAfter(idx uint, node *Node[T]) {
  old_node := l.get(idx)
  temp := old_node.Next
  old_node.Next = node
  node.Next = temp
  l.length++
}

func (l *LinkedList[T]) removeAt(idx uint) *Node[T] {
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




