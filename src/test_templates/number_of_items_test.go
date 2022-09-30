package kata_tests

import (
  "testing"
  "github.com/geefuoco/go-kata-machine/katas"
  "github.com/stretchr/testify/assert"
)

func TestNumberOfItems(t *testing.T){
  s := "*|**|*|"
  start := []int{1, 1}
  end := []int{5, 7}
  assert.Equal(t, []int{2, 3}, katas.NumberOfItems(s, start, end))
}
