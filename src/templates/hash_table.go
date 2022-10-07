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

func (h *HashTable[K, V]) Insert(key string, value V) {
}

func (h *HashTable[K, V]) Has(key string) bool {
  return false
}

func (h *HashTable[K, V]) Get(key string) V{
  var out V
  return out
}

func (h *HashTable[K, V]) Delete(key string) V{
  var out V
  return out
}

