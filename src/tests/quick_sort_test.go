package tests

import (
  "testing"
  "github.com/geefuoco/go-kata-machine/util"
  "github.com/stretchr/testify/assert"
)

func TestQuickSort(t *testing.T){

  arr := []int{9, 5, 4, 8, 2, 3, 1, 6, 7}
  util.QuickSort(arr)
  assert.Equal(t, []int{1, 2, 3, 4, 5, 6, 7,8, 9}, arr)

  arr2 := []int{3, 3, 3, 9, 9, 9, 1, 2, 3, 0}
  util.QuickSort(arr2)
  assert.Equal(t, []int{0, 1, 2, 3, 3, 3, 3,9, 9, 9}, arr2)
}
