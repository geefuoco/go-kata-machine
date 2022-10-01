package tests

import (
  "testing"
  "github.com/geefuoco/go-kata-machine/src/util"
  "github.com/stretchr/testify/assert"
)

func TestLargestComponent(t *testing.T){
  graph := map[rune][]rune{
    'a': {'b', 'c'},
    'b': {'a'},
    'c': {'a'},
    'd': {'e'},
    'e': {'d', 'g', 'f'},
    'f': {'e'},
    'g': {'e'},
  }

  assert.Equal(t, 4, util.LargestComponent(graph))
}
