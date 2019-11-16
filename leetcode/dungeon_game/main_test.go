package dungeon_game

import "testing"

func TestBasic(t *testing.T) {
	// -6 is the smallest negative which will be the minimum positive.
	//calculateMinimumHP([][]int{
	//	{-2, -3, 3},
	//	{-5, -10, 1},
	//	{10, 30, -5},
	//})

	calculateMinimumHP([][]int{
		{1, 1, 1},
		{-1, 0, 1},
		{-3, 0, 2},
	})
}
