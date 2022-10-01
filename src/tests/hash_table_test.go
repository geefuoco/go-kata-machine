package tests

import (
  "testing"
  "github.com/geefuoco/go-kata-machine/src/util"
  "github.com/stretchr/testify/assert"
)

func TestHashTable(t *testing.T){
  hash := util.NewHashTable[string]()
  hash.Insert("mario", "kart")
  hash.Insert("tokyo", "drift")
  assert.Equal(t, "kart", hash.Get("mario"))
  assert.Equal(t, "drift", hash.Get("tokyo"))
  hash.Delete("mario")
  assert.Equal(t, "", hash.Get("mario"))
  hash.Delete("tokyo")
  assert.Equal(t, "", hash.Get("tokyo"))
  nums := util.NewHashTable[int]()
  nums.Insert("one", 1)
  nums.Insert("two", 2)
  assert.Equal(t, 1, nums.Get("one"))
  assert.Equal(t, 2, nums.Get("two"))
  assert.Equal(t, false, nums.Has("twenty"))
}
