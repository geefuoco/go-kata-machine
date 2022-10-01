package tests

import (
  "testing"
  "github.com/geefuoco/go-kata-machine/src/util"
  "github.com/stretchr/testify/assert"
)

func TestHasCycle(t *testing.T) {
  graph := map[rune][]rune{
    'a': {'b', 'c'},
    'b': {'d', 'f'},
    'c': {'a', 'e'},
  }

  assert.Equal(t, util.HasCycle(graph), true)

  acyclic := map[rune][]rune{
    'a': {'b', 'c'},
    'b': {'d', 'e'},
  }
  assert.Equal(t, util.HasCycle(acyclic), false)
  
}
