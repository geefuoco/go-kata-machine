package tests

import (
  "testing"
  "github.com/geefuoco/go-kata-machine/util"
  "github.com/stretchr/testify/assert"
)

func TestMaxHeap(t *testing.T) {

  m := &util.MaxHeap{}
  m.Push(10)
  m.Push(11)
  m.Push(7)
  m.Push(100)
  v := m.Pop()
  assert.Equal(t, 100, v)
  assert.Equal(t, 11, m.Pop())
  m.Push(19)
  assert.Equal(t, 19, m.Pop())
  assert.Equal(t, 10, m.Pop())
}

