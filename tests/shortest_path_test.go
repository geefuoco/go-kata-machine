package tests

import (
  "testing"
  "github.com/geefuoco/go-kata-machine/util"
  "github.com/stretchr/testify/assert"
)

func TestShortestPath(t *testing.T){
  edges := [][2]rune{
    {'w', 'x'},
    {'x', 'y'},
    {'z', 'y'},
    {'z', 'v'},
    {'w', 'v'},
  }

  assert.Equal(t, 2, util.ShortestPath(edges, 'w', 'y'))
  assert.Equal(t, 2, util.ShortestPath(edges, 'w', 'z'))

  longEdges := [][2]rune{
    {'a', 'b'},
    {'b', 'c'},
    {'c', 'd'},
    {'d', 'e'},
    {'e', 'f'},
    {'f', 'g'},
    {'g', 'h'},
    {'h', 'i'},
    {'i', 'j'},
    {'j', 'k'},
    {'k', 'a'},
  }

  assert.Equal(t, 5, util.ShortestPath(longEdges, 'a', 'g'))
}
