package util

import "reflect"

const ArraySize = 50

type Key interface {
  string | int
}


type HashTable[K Key, V any] struct {
  array [ArraySize]*bucket[K, V]
}

type bucketNode[K Key, V any] struct {
  key K
  val V
  next *bucketNode[K, V]
}

type bucket[K Key, V any] struct {
  head *bucketNode[K, V]
}

func (b*bucket[K, V]) has(key K) bool {
  current := b.head
  for current != nil {
    if current.key == key {
      return true
    }
    current = current.next
  }
  return false
}

func (b*bucket[K, V]) get(key K) *bucketNode[K, V] {
  current := b.head
  for current != nil {
    if current.key == key {
      return current
    }
    current = current.next
  }
  return current
}

func (b*bucket[K, V]) insert(key K, val V) {
  var node *bucketNode[K, V]
  if b.has(key) {
    node = b.get(key)
    node.val = val
  } else {
    node = &bucketNode[K, V]{key, val, nil}
    next := b.head
    b.head = node
    node.next = next
  }
}

func (b*bucket[K, V]) remove(key K) *bucketNode[K, V] {
  var prev *bucketNode[K, V]
  current := b.head
  for current != nil {
    if current.key == key {
      if prev != nil {
        prev.next = current.next
      } else {
        b.head = nil
      }
      return current
    }
    prev = current
    current = current.next
  }
  return nil
}

func NewHashTable[K Key, V any]() *HashTable[K, V]{
  return &HashTable[K, V]{array: [ArraySize]*bucket[K, V]{}}
}

func (h *HashTable[K, V]) Insert(key K, value V) {
  hashed := hash(key)
  bin := h.array[hashed]
  if bin == nil {
    bin = &bucket[K, V]{}
    h.array[hashed] = bin
  }
  bin.insert(key, value)
}

func (h *HashTable[K, V]) Has(key K) bool {
  hashed := hash(key)
  bin := h.array[hashed]
  if bin == nil {
    return false
  }
  return bin.has(key)
}

func (h *HashTable[K, V]) Get(key K) V{
  var out V
  hashed := hash(key)
  bin := h.array[hashed]
  if bin != nil{
    if node := bin.get(key); node != nil {
      out = node.val
    }
  }
  return out
}

func (h *HashTable[K, V]) Delete(key K) V{
  var out V
  hashed := hash(key)
  bin := h.array[hashed]
  if bin != nil{
    if node := bin.remove(key); node != nil {
      out = node.val
    }
  }
  return out
}

func hash[K Key](key K) int {
  if reflect.TypeOf(key) == reflect.TypeOf("hello") {
    str, ok := any(key).(string)
    if !ok{
      panic("Error while hashing string key")
    }
    var sum int
    for _, char := range str {
      sum += int(char)
    }
    return sum % ArraySize
  }
  val, ok := any(key).(int)
  if !ok {
    panic("Error while hashing int key")
  }
  return val % ArraySize
}
