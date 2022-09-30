package katas

type Stack[T any] struct {
}

func NewStack[T any]() *Stack[T] {
}


func (s *Stack[T]) Push(val T) {
}

func (s *Stack[T]) Pop() T {
}

func (s *Stack[T]) Peak() T {
}

func (s *Stack[T]) IsEmpty() bool {
}

func (s *Stack[T]) Length() int {
}
