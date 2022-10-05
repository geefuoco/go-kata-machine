package kata_tests

import (
  "testing"
  "github.com/geefuoco/go-kata-machine/katas"
  "github.com/stretchr/testify/assert"
)

func TestHasCycle(t *testing.T) {
  graph := map[rune][]rune{
    'a': {'b', 'c'},
    'b': {'d', 'f'},
    'c': {'a', 'e'},
  }

  assert.Equal(t, true, katas.HasCycle(graph))

  acyclic := map[rune][]rune{
    'a': {'b', 'c'},
    'b': {'d', 'e'},
  }
  assert.Equal(t, false, katas.HasCycle(acyclic))
  
}
