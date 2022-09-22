package util

type Stack[T any] struct {
  head *Node[T]
  length int
}

func NewStack[T any]() *Stack[T] {
  return &Stack[T]{head: nil, length: 0}
}


func (s *Stack[T]) Push(val T) {
  node := Node[T]{Val: val}
  temp := s.head
  s.head = &node
  node.Next = temp
  s.length++
}

func (s *Stack[T]) Pop() T {
  var val T
  if s.head == nil {
    return val
  }
  out := s.head
  s.head = out.Next
  if s.length < 1 {
    s.length = 0
  } else {
    s.length--
  }
  return out.Val
}

func (s *Stack[T]) Peak() T {
  var val T
  if s.head == nil {
    return val
  }
  return s.head.Val
}

func (s *Stack[T]) Empty() {
  s.length = 0
  s.head = nil
}

func (s *Stack[T]) Length() int {
  return s.length
}
