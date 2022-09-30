package tests

import (
  "testing"
  "github.com/geefuoco/go-kata-machine/util"
  "github.com/stretchr/testify/assert"
)

func TestStack(t *testing.T) {

  stack := util.NewStack[int]()
  stack.Push(15)
  assert.Equal(t, 1, stack.Length())
  stack.Push(10)
  assert.Equal(t, 2, stack.Length())
  assert.Equal(t, 10, stack.Peak())
  out := stack.Pop()
  assert.Equal(t, 10, out)
  assert.Equal(t, 1, stack.Length())
  assert.Equal(t, 15, stack.Peak())
  stack.Pop()
  assert.Equal(t, 0, stack.Length())
  stack.Pop()
  assert.Equal(t, 0, stack.Length())

}
