package katas


type Node[T any] struct {
  Val T
  Next *Node[T]
  Prev *Node[T]
}
