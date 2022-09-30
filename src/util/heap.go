package util

type MaxHeap struct {
  array []int
}

func (h *MaxHeap) Push(value int) {
  h.array = append(h.array, value)
  h.heapifyUp(len(h.array)-1)
}

func (h *MaxHeap) Pop() int {
  if len(h.array) < 1{
    return -1
  }
  out := h.array[0]
  h.array[0] = h.array[len(h.array)-1]
  h.array = h.array[:len(h.array)-1]
  h.heapifyDown(0)
  return out
}

func (h *MaxHeap) heapifyUp(index int) {
  for h.array[parent(index)] < h.array[index] {
    h.swap(parent(index), index)
    index=  parent(index)
  }
}

func (h *MaxHeap) heapifyDown(index int) {
  last := len(h.array)-1
  l, r:= left(index), right(index)
  compare := 0
  for l <= last {
    if l == last || h.array[l] > h.array[r] {
      compare = l
    } else {
      compare = r
    }

    if h.array[index] < h.array[compare]{
      h.swap(index, compare)
      index = compare
      l, r = left(index), right(index)
    } else {
      return
    }
  }
}

func parent(i int) int {
  return (i-1)/2
}

func left(i int) int {
  return 2*i + 1
}

func right(i int) int{
  return 2*i + 2
}

func (h *MaxHeap) swap(i1, i2 int) {
  h.array[i1], h.array[i2] = h.array[i2], h.array[i1]
}

