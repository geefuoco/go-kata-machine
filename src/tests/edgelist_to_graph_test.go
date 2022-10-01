package tests

import (
  "testing"
  "github.com/stretchr/testify/assert"
  "github.com/geefuoco/go-kata-machine/src/util"
)

func TestConvertEdgesToGraph(t *testing.T){
  edges := [][2]rune{
    {'w', 'x'},
    {'x', 'y'},
    {'z', 'y'},
    {'z', 'v'},
    {'w', 'v'},
  }

  adjacent := make(util.Graph)
  adjacent['w']=[]rune{'x', 'v'}
  adjacent['x']=[]rune{'w', 'y'}
  adjacent['z']=[]rune{'y', 'v'}
  adjacent['y']=[]rune{'x', 'z'}
  adjacent['v']=[]rune{'z', 'w'}

  assert.Equal(t, adjacent, util.ConvertEdgelistToGraph(edges))
}
