package katas

const ArraySize = 50

type Key interface {
  string | int
}


type HashTable[K Key, V any] struct {
}


func NewHashTable[K Key, V any]() *HashTable[K, V]{
  return nil
}

func (h *HashTable[K, V]) Insert(key K, value V) {
}

func (h *HashTable[K, V]) Has(key K) bool {
  return false
}

func (h *HashTable[K, V]) Get(key K) V{
  var out V
  return out
}

func (h *HashTable[K, V]) Delete(key K) V{
  var out V
  return out
}

