package kata_tests

import (
  "testing"
  "github.com/geefuoco/go-kata-machine/katas"
  "github.com/stretchr/testify/assert"
)

func TestMergeSort(t *testing.T){

  arr := []int{9, 5, 4, 8, 2, 3, 1, 6, 7}
  katas.MergeSort(arr)
  assert.Equal(t, []int{1, 2, 3, 4, 5, 6, 7,8, 9}, arr)

  arr2 := []int{2, 1, 3, 5, 6, 4, 7}
  katas.MergeSort(arr2)
  assert.Equal(t, []int{1, 2, 3, 4, 5, 6, 7}, arr2)
  arr3 := []int{1,2, 3, 5, 6, 4, 7}
  katas.MergeSort(arr3)
  assert.Equal(t, []int{1, 2, 3, 4, 5, 6, 7}, arr3)

}
