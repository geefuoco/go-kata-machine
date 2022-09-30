package tests

import (
  "testing"
  "github.com/geefuoco/go-kata-machine/util"
  "github.com/stretchr/testify/assert"
)

func TestTrie(t *testing.T) {
  trie := util.NewTrie()
  trie.Insert("hello")
  assert.Equal(t, trie.Contains("hello"), true)
  assert.Equal(t, trie.Contains("hell"), false)
  assert.Equal(t, trie.Contains("petunia"), false)
  trie.Insert("hell")
  assert.Equal(t, trie.Contains("hell"), true)
}
