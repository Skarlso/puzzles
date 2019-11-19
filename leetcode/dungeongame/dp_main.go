package dungeongame

import "math"

func calculateMinimumHPDP(dungeon [][]int) int {
	n := len(dungeon)
	m := len(dungeon[0])

	dp := make([][]int, n+1)
	for i := 0; i < n+1; i++ {
		dp[i] = make([]int, m+1)
	}

	for i := 0; i < n+1; i++ {
		for j := 0; j < m+1; j++ {
			dp[i][j] = math.MaxInt64
		}
	}
	dp[n][m-1] = 1
	dp[n-1][m] = 1
	for i := n - 1; i >= 0; i-- {
		for j := m - 1; j >= 0; j-- {
			dp[i][j] = max(1, min(dp[i+1][j], dp[i][j+1])-dungeon[i][j])
		}
	}

	return dp[0][0]
}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}

func min(a, b int) int {
	if a < b {
		return a
	}
	return b
}
