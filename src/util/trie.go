package util

type TrieNode struct  {
  Children map[rune]*TrieNode
  Terminate bool
}

type Trie struct {
  Root *TrieNode
}

func NewTrie() *Trie {
  return &Trie{Root: &TrieNode{Children: make(map[rune]*TrieNode)}}
}

func (t *Trie) Insert(word string) {
  node := t.Root
  for _, c := range word {
    next := node.Children[c]
    if next == nil {
      node.Children[c] = &TrieNode{Children: make(map[rune]*TrieNode)}
      next = node.Children[c]
    }
    node = next
  }
  node.Terminate = true
}

func (t *Trie) Contains(word string) bool {
  node := t.Root
  for _, c := range word {
    next := node.Children[c]
    if next == nil {
      return false
    }
    node = next
  }
  return node.Terminate
}

