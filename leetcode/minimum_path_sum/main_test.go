package main

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestBasic(t *testing.T) {
	/*
		1, 3, 1
		1, 5, 1
		4, 2, 1
	*/
	grid := [][]int{{1, 3, 1}, {1, 5, 1}, {4, 2, 1}}
	expect := 7

	got := minPathSum(grid)
	assert.Equal(t, expect, got, "minimum should have been 7")
}

func Test2Basic(t *testing.T) {
	/*
		1, 2
		1, 1
	*/
	grid := [][]int{{1, 2}, {1, 1}}
	expect := 3

	got := minPathSum(grid)
	assert.Equal(t, expect, got, "minimum should have been 3")
}
