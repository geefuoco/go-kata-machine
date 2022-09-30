package kata_tests

import (
  "testing"
  "github.com/geefuoco/go-kata-machine/katas"
  "github.com/stretchr/testify/assert"
)

func TestTrie(t *testing.T) {
  trie := katas.NewTrie()
  trie.Insert("hello")
  assert.Equal(t, trie.Contains("hello"), true)
  assert.Equal(t, trie.Contains("hell"), false)
  assert.Equal(t, trie.Contains("petunia"), false)
  trie.Insert("hell")
  assert.Equal(t, trie.Contains("hell"), true)
}
