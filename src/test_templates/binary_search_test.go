package kata_tests

import (
  "testing"
  "github.com/geefuoco/go-kata-machine/katas"
  "github.com/stretchr/testify/assert"
)

func TestBinarySearch(t *testing.T) {
  array := []int{1, 2, 3, 4, 5, 6, 7, 8, 9}

  assert.Equal(t, true, katas.BinarySearch(array, 1))
  assert.Equal(t, true, katas.BinarySearch(array, 2))
  assert.Equal(t, true, katas.BinarySearch(array, 3))
  assert.Equal(t, true, katas.BinarySearch(array, 4))
  assert.Equal(t, true, katas.BinarySearch(array, 5))
  assert.Equal(t, true, katas.BinarySearch(array, 6))
  assert.Equal(t, false, katas.BinarySearch(array, 10))

}
