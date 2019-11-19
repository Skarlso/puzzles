package cherry_pickup

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestBasic(t *testing.T) {
	grid := [][]int{
		{0, 1, -1},
		{1, 0, -1},
		{1, 1, 1},
	}
	s := cherryPickup(grid)
	assert.Equal(t, 5, s, "should equal")
}

// [[1,1,-1],[1,-1,1],[-1,1,1]]
func Test2Basic(t *testing.T) {
	grid := [][]int{
		{1, 1, -1},
		{1, -1, 1},
		{-1, 1, 1},
	}
	s := cherryPickup(grid)
	assert.Equal(t, 0, s, "should equal")
}

func Test3Basic(t *testing.T) {
	grid := [][]int{
		{1},
	}
	s := cherryPickup(grid)
	assert.Equal(t, 1, s, "should equal")
}

// [[1,1,1,1,0,0,0],[0,0,0,1,0,0,0],[0,0,0,1,0,0,1],[1,0,0,1,0,0,0],[0,0,0,1,0,0,0],[0,0,0,1,0,0,0],[0,0,0,1,1,1,1]]
// because first is ----|||----|||||| --> It takes a less optimal path in the first run so the second run
// can pick up all of the remaining cherries.
func Test4Basic(t *testing.T) {
	grid := [][]int{
		{1, 1, 1, 1, 0, 0, 0},
		{0, 0, 0, 1, 0, 0, 0},
		{0, 0, 0, 1, 0, 0, 1},
		{1, 0, 0, 1, 0, 0, 0},
		{0, 0, 0, 1, 0, 0, 0},
		{0, 0, 0, 1, 0, 0, 0},
		{0, 0, 0, 1, 1, 1, 1},
	}
	s := cherryPickup(grid)
	assert.Equal(t, 15, s, "should equal")
}
