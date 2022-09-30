package util

const ArraySize = 50

type Key interface {
  string | int
}


type HashTable[V any] struct {
  array [ArraySize]*bucket[V]

}

type bucket[V any] struct {
  head *bucketNode[V]
}

type bucketNode[V any] struct {
  key string
  value V
  next *bucketNode[V]
}

func NewHashTable[V any]() *HashTable[V]{
  table := &HashTable[V]{}
  for i:= range table.array{
    table.array[i] = &bucket[V]{}
  }
  return table
}

func (h *HashTable[V]) Insert(key string, value V) {
  index := hash(key)
  bucket := h.array[index]
  bucket.insert(key, value)
}

func (h *HashTable[V]) Has(key string) bool {
  index := hash(key)
  return h.array[index].has(key)
}

func (h *HashTable[V]) Get(key string) V{
  index := hash(key)
  return h.array[index].search(key)
}

func (h *HashTable[V]) Delete(key string) V{
  index :=hash(key)
  return h.array[index].delete(key)
}

func (b *bucket[V]) insert(key string, value V){
  node := &bucketNode[V]{key: key, value: value}
  temp := b.head
  b.head = node
  b.head.next = temp
}

func (b *bucket[V]) search(key string) V {
  current := b.head
  var value V
  for current != nil {
    if current.key == key{
      value = current.value
      break
    }
    current = current.next
  }
  return value 
}

func (b *bucket[V]) has(key string) bool {
  current := b.head
  for current != nil {
    if current.key == key{
      return true
    }
    current = current.next
  }
  return false 
}

func (b *bucket[V]) delete(key string) V {
  current := b.head
  var prev *bucketNode[V]
  var value V
  for current != nil {
    if current.key == key{
      if prev != nil {
        prev.next = current.next
      } else {
        b.head = nil
      }
      value = current.value
      break
    }
    prev = current
    current = current.next
  }
  return value
}

func hash(key string) int {
  var sum int
  for _, c := range key {
    sum += int(c)
  }
  return sum%ArraySize
}
