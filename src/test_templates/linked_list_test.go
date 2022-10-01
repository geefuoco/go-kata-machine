package kata_tests

import (
  "testing"
  "github.com/stretchr/testify/assert"
  "github.com/geefuoco/go-kata-machine/katas"
)

type Node = katas.Node[int];

func TestLinkedList(t *testing.T) {
  list := katas.NewLinkedList[int]()
  assert.Equal(t, list.Length(), 0)

  list.Append(15)
  list.Prepend(19)
  assert.Equal(t, 19, list.Head().Val)
  assert.Equal(t, 15, list.Tail().Val)

  list.Append(20)
  assert.Equal(t, 3, list.Length())
  list.Append(25)
  assert.Equal(t, 4, list.Length())
  list.Append(39)
  assert.Equal(t, 5, list.Length())
  list.Append(81)
  assert.Equal(t, 6, list.Length())
  list.Append(1)
  assert.Equal(t, 7, list.Length())
  assert.Equal(t, 1, list.Tail().Val)
  assert.Equal(t, 19, list.Get(0))
  assert.Equal(t, 1, list.Get(list.Length()-1))
  list.InsertAfter(3, 4)
  assert.Equal(t, 4, list.Get(4))
  assert.Equal(t, 8, list.Length())
  out := list.RemoveAt(7)
  assert.Equal(t, 1, out)
  assert.Equal(t, 7, list.Length())
  h := list.RemoveAt(0)
  assert.Equal(t, 19, h)
  assert.Equal(t, 15, list.Head().Val)
  assert.Equal(t, 6, list.Length())

  list.Pop()
  assert.Equal(t, 5, list.Length())
  list.Pop()
  assert.Equal(t, 4, list.Length())
  list.Pop()
  assert.Equal(t, 3, list.Length())
  list.Pop()
  assert.Equal(t, 2, list.Length())
  list.Pop()
  assert.Equal(t, 1, list.Length())
  list.Pop()
  assert.Equal(t, 0, list.Length())

  assert.Equal(t, (*Node)(nil), list.Head())
  assert.Equal(t, (*Node)(nil), list.Tail())

}

func TestPanic(t *testing.T) {
  defer func() {
    if r := recover(); r == nil {
      t.Errorf("Test did not panic")
    }
  }()
  list := katas.NewLinkedList[int]()
  list.Get(100)
}

