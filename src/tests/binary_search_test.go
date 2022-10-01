package tests

import (
  "testing"
  "github.com/geefuoco/go-kata-machine/src/util"
  "github.com/stretchr/testify/assert"
)

func TestBinarySearch(t *testing.T) {
  array := []int{1, 2, 3, 4, 5, 6, 7, 8, 9}

  assert.Equal(t, true, util.BinarySearch(array, 1))
  assert.Equal(t, true, util.BinarySearch(array, 2))
  assert.Equal(t, true, util.BinarySearch(array, 3))
  assert.Equal(t, true, util.BinarySearch(array, 4))
  assert.Equal(t, true, util.BinarySearch(array, 5))
  assert.Equal(t, true, util.BinarySearch(array, 6))
  assert.Equal(t, false, util.BinarySearch(array, 10))

}
