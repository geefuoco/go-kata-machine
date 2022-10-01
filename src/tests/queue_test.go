package tests

import (
  "testing"
  "github.com/geefuoco/go-kata-machine/src/util"
  "github.com/stretchr/testify/assert"
)

func TestQueue(t *testing.T) {
  q := util.NewQueue[int]()
  q.Enqueue(7)
  assert.Equal(t, 7, q.Peak())
  q.Deque()
  assert.Equal(t, 0, q.Peak())
  assert.Equal(t, 0, q.Length())
  assert.Equal(t, (*util.Node[int])(nil), q.Head())
  assert.Equal(t, (*util.Node[int])(nil), q.Tail())
  q.Deque()
  assert.Equal(t, 0, q.Length())
  q.Enqueue(13)
  q.Enqueue(14)
  assert.Equal(t, 14, q.Tail().Val)
  assert.Equal(t, 2, q.Length())
  assert.Equal(t, 13, q.Deque())
  assert.Equal(t, 14, q.Deque())
}
