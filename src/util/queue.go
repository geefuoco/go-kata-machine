package util

type Queue[T any] struct {
  head *Node[T]
  tail *Node[T]
  length int
}

func NewQueue[T any]() *Queue[T] {
  return &Queue[T]{}
}

func (q *Queue[T])Enqueue(val T) {
  q.length++
  node := &Node[T]{Val: val}
  if q.head == nil {
    q.head = node
    q.tail = node
  } else {
    q.tail.Next = node
    q.tail = node
  }
}

func (q *Queue[T])Deque() T {
  if q.length > 0 {
    q.length--
  }
  var out T
  if q.head != nil {
    out = q.head.Val
    q.head = q.head.Next
    if q.head == nil {
      q.tail = nil
    }
  }
  return out
}

func (q *Queue[T])Peak() T {
  var out T
  if q.head != nil {
    return q.head.Val
  }
  return out
}

func (q *Queue[T])Length() int {
  return q.length
}

func (q *Queue[T])Head() *Node[T] {
  return q.head
}

func (q *Queue[T])Tail() *Node[T] {
  return q.tail
}
