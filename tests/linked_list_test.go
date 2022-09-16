package tests

import (
  "testing"
  "github.com/stretchr/testify/assert"
  "github.com/geefuoco/go-kata-machine/util"
)

type Node = util.Node[uint];

func TestLinkedList(t *testing.T) {
  list := util.NewLinkedList[uint]()
  assert.Equal(t, list.Length(), uint(0))
  node := &Node{Val: 15, Next:nil, Prev: nil}

  list.Append(node)

  assert.Equal(t, list.Head(), node)
  assert.Equal(t, list.Tail(), node)

  list.Prepend(&Node{Val: 19, Next: nil, Prev: nil})
  assert.Equal(t, list.Head().Val, uint(19))
  assert.Equal(t, list.Tail().Val, uint(15))

  list.Append(&Node{Val: 20, Next: nil, Prev: nil})
  assert.Equal(t, list.Length(), uint(3))
  list.Append(&Node{Val: 25, Next: nil, Prev: nil})
  assert.Equal(t, list.Length(), uint(4))
  list.Append(&Node{Val: 39, Next: nil, Prev: nil})
  assert.Equal(t, list.Length(), uint(5))
  list.Append(&Node{Val: 81, Next: nil, Prev: nil})
  assert.Equal(t, list.Length(), uint(6))
  list.Append(&Node{Val: 1, Next: nil, Prev: nil})
  assert.Equal(t, list.Length(), uint(7))
  assert.Equal(t, list.Tail().Val, uint(1))
  assert.Equal(t, list.Get(0).Val, uint(19))
  assert.Equal(t, list.Get(list.Length()-1).Val, uint(1))
  list.InsertAfter(3, &Node{Val: 4, Next: nil, Prev: nil})
  assert.Equal(t, list.Get(4).Val, uint(4))
  assert.Equal(t, list.Length(), uint(8))
  out := list.RemoveAt(7)
  assert.Equal(t, out.Val, uint(1))
  assert.Equal(t, list.Length(), uint(7))
  h := list.RemoveAt(0)
  assert.Equal(t, h.Val, uint(19))
  assert.Equal(t, list.Head().Val, uint(15))
  assert.Equal(t, list.Length(), uint(6))

  list.Pop()
  assert.Equal(t, list.Length(), uint(5))
  list.Pop()
  assert.Equal(t, list.Length(), uint(4))
  list.Pop()
  assert.Equal(t, list.Length(), uint(3))
  list.Pop()
  assert.Equal(t, list.Length(), uint(2))
  list.Pop()
  assert.Equal(t, list.Length(), uint(1))
  list.Pop()
  assert.Equal(t, list.Length(), uint(0))

  assert.Equal(t, list.Head(), (*Node)(nil))
  assert.Equal(t, list.Tail(), (*Node)(nil))

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

