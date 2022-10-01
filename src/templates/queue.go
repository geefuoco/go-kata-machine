package katas

type Queue[T any] struct {
}

func NewQueue[T any]() *Queue[T] {
  return nil
}

func (q *Queue[T])Enqueue(val T) {
}

func (q *Queue[T])Deque() T {
  var out T
  return out
}

func (q *Queue[T])Peak() T {
  var out T
  return out
}

func (q *Queue[T])IsEmpty() bool {
  return false
}

func (q *Queue[T])Length() int {
  return -1
}

func (q *Queue[T])Head() *Node[T] {
  return nil
}

func (q *Queue[T])Tail() *Node[T] {
  return nil
}
