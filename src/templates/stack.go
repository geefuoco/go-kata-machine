package katas

type Stack[T any] struct {
}

func NewStack[T any]() *Stack[T] {
  return nil
}


func (s *Stack[T]) Push(val T) {
}

func (s *Stack[T]) Pop() T {
  var out T
  return out
}

func (s *Stack[T]) Peak() T {
  var out T
  return out
}

func (s *Stack[T]) IsEmpty() bool {
  return false
}

func (s *Stack[T]) Length() int {
  return -1
}
