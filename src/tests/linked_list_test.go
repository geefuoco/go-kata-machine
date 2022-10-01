package tests

import (
  "testing"
  "github.com/stretchr/testify/assert"
  "github.com/geefuoco/go-kata-machine/src/util"
)

type Node = util.Node[uint];

func TestLinkedList(t *testing.T) {
  list := util.NewLinkedList[uint]()
  assert.Equal(t, list.Length(), uint(0))

  list.Append(15)
  list.Prepend(19)
  assert.Equal(t, uint(19), list.Head().Val)
  assert.Equal(t, uint(15), list.Tail().Val)

  list.Append(20)
  assert.Equal(t, uint(3), list.Length())
  list.Append(25)
  assert.Equal(t, uint(4), list.Length())
  list.Append(39)
  assert.Equal(t, uint(5), list.Length())
  list.Append(81)
  assert.Equal(t, uint(6), list.Length())
  list.Append(1)
  assert.Equal(t, uint(7), list.Length())
  assert.Equal(t, uint(1), list.Tail().Val)
  assert.Equal(t, uint(19), list.Get(0))
  assert.Equal(t, uint(1), list.Get(list.Length()-1))
  list.InsertAfter(3, 4)
  assert.Equal(t, uint(4), list.Get(4))
  assert.Equal(t, uint(8), list.Length())
  out := list.RemoveAt(7)
  assert.Equal(t, uint(1), out)
  assert.Equal(t, uint(7), list.Length())
  h := list.RemoveAt(0)
  assert.Equal(t, uint(19), h)
  assert.Equal(t, uint(15), list.Head().Val)
  assert.Equal(t, uint(6), list.Length())

  list.Pop()
  assert.Equal(t, uint(5), list.Length())
  list.Pop()
  assert.Equal(t, uint(4), list.Length())
  list.Pop()
  assert.Equal(t, uint(3), list.Length())
  list.Pop()
  assert.Equal(t, uint(2), list.Length())
  list.Pop()
  assert.Equal(t, uint(1), list.Length())
  list.Pop()
  assert.Equal(t, uint(0), list.Length())

  assert.Equal(t, (*Node)(nil), list.Head())
  assert.Equal(t, (*Node)(nil), list.Tail())

}

func TestPanic(t *testing.T) {
  defer func() {
    if r := recover(); r == nil {
      t.Errorf("Test did not panic")
    }
  }()
  list := util.NewLinkedList[uint]()
  list.Get(100)
}

