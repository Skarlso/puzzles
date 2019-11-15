package main

func minPathSum(grid [][]int) int {
	ly := len(grid)
	lx := len(grid[0])
	dp := make([][]int, ly)
	for i := 0; i < ly; i++ {
		dp[i] = make([]int, lx)
	}

	dp[0][0] = grid[0][0]

	// update the first row
	for i := 1; i < lx; i++ {
		dp[0][i] = dp[0][i - 1] + grid[0][i]
	}

	// update the first column
	for i := 1; i < ly; i++ {
		dp[i][0] = dp[i - 1][0] + grid[i][0]
	}

	for i := 1; i < ly; i++ {
		for j := 1; j < lx; j++ {
			dp[i][j] = min(dp[i - 1][j], dp[i][j - 1]) + grid[i][j]
		}
	}

	return dp[ly - 1][lx - 1]
}

func min(a, b int) int {
	if a < b {
		return a
	}
	return b
}
