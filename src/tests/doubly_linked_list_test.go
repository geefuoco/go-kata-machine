package tests


import (
  "testing"
  "github.com/geefuoco/go-kata-machine/util"
  "github.com/stretchr/testify/assert"
)

func TestDoublyLinkedList(t *testing.T) {
  d := util.NewDoublyLinkedList[int]()

  assert.Equal(t, 0, d.Length())
  d.Append(1)
  assert.Equal(t, 1, d.Length())
  d.Pop()
  assert.Equal(t, 0, d.Length())
  assert.Equal(t, (*util.Node[int])(nil), d.Head())
  assert.Equal(t, (*util.Node[int])(nil), d.Tail())
  d.Prepend(14)
  d.Append(50)
  d.InsertAt(1, 30)
  assert.Equal(t, 3, d.Length())
  assert.Equal(t, 30, d.Get(1))
  assert.Equal(t, 14, d.Get(0))
  assert.Equal(t, 50, d.Get(2))
  d.InsertAt(1, 40)
  assert.Equal(t, 40, d.Get(1))
  d.InsertAt(1, 60)
  assert.Equal(t, 60, d.Get(1))
  d.InsertAt(1, 70)
  assert.Equal(t, 70, d.Get(1))
  d.InsertAt(1, 80)
  assert.Equal(t, 80, d.Get(1))
  assert.Equal(t, 7, d.Length())
}

