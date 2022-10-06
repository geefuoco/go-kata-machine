package kata_tests

import (
  "testing"
  "github.com/geefuoco/go-kata-machine/katas"
  "github.com/stretchr/testify/assert"
)

func TestHashTable(t *testing.T){
  hash := katas.NewHashTable[string, string]()
  hash.Insert("mario", "kart")
  hash.Insert("tokyo", "drift")
  assert.Equal(t, "kart", hash.Get("mario"))
  assert.Equal(t, "drift", hash.Get("tokyo"))
  hash.Delete("mario")
  assert.Equal(t, "", hash.Get("mario"))
  hash.Delete("tokyo")
  assert.Equal(t, "", hash.Get("tokyo"))
  nums := katas.NewHashTable[int, int]()
  nums.Insert(10, 1)
  nums.Insert(20, 2)
  assert.Equal(t, 1, nums.Get(10))
  assert.Equal(t, 2, nums.Get(20))
  assert.Equal(t, false, nums.Has(30))
}
