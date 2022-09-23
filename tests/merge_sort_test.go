package tests

import (
  "testing"
  "github.com/geefuoco/go-kata-machine/util"
  "github.com/stretchr/testify/assert"
)

func TestMergeSort(t *testing.T){

  arr := []int{9, 5, 4, 8, 2, 3, 1, 6, 7}
  util.MergeSort(arr)
  assert.ElementsMatch(t, []int{1, 2, 3, 4, 5, 6, 7,8, 9}, arr)
}
