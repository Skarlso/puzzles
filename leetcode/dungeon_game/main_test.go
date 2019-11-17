package dungeon_game

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestBasic(t *testing.T) {
	minLife := calculateMinimumHP([][]int{
		{-2, -3, 3},
		{-5, -10, 1},
		{10, 30, -5},
	})
	assert.Equal(t, 7, minLife, "minimum life should equal the given value")
}

func Test2Basic(t *testing.T) {
	minLife := calculateMinimumHP([][]int{
		{1, -3, 2},
		{0, -1, 2},
		{0, 0, -2},
	})
	assert.Equal(t, 1, minLife, "minimum life should equal the given value")
}

func Test3Basic(t *testing.T) {
	minLife := calculateMinimumHP([][]int{
		{-3, 5},
	})
	assert.Equal(t, 4, minLife, "if the knights life begins with -number he immediately dies")
}
