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
