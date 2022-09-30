package katas

type Queue[T any] struct {
}

func NewQueue[T any]() *Queue[T] {
}

func (q *Queue[T])Enqueue(val T) {
}

func (q *Queue[T])Deque() T {
}

func (q *Queue[T])Peak() T {
}

func (q *Queue[T])IsEmpty() bool {
}

func (q *Queue[T])Length() int {
}

func (q *Queue[T])Head() *Node[T] {
}

func (q *Queue[T])Tail() *Node[T] {
}
