package kata_tests

import (
  "testing"
  "github.com/geefuoco/go-kata-machine/katas"
  "github.com/stretchr/testify/assert"
)

func TestServedBuildings(t *testing.T){
  buildings := []int{1, 2, 1, 2, 2}
  routerLocations := []int{3, 1}
  routerRanges := []int{1, 2}

  assert.Equal(t, 3, katas.ServedBuildings(buildings, routerLocations, routerRanges))

  buildings2 := []int{1, 2, 2, 1, 2, 1, 1, 2, 2, 5, 6, 2}
  routerLocations2 := []int{1, 6, 9}
  routerRanges2 := []int{1, 6, 2}
  assert.Equal(t, 7, katas.ServedBuildings(buildings2, routerLocations2, routerRanges2))
}
