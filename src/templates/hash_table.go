package katas

const ArraySize = 50

type Key interface {
  string | int
}


type HashTable[V any] struct {
}


func NewHashTable[V any]() *HashTable[V]{
  return nil
}

func (h *HashTable[V]) Insert(key string, value V) {
}

func (h *HashTable[V]) Has(key string) bool {
  return false
}

func (h *HashTable[V]) Get(key string) V{
  var out V
  return out
}

func (h *HashTable[V]) Delete(key string) V{
  var out V
  return out
}

