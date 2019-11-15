package main

func uniquePaths(m int, n int) int {
	path := 1
	for i := n; i < (m + n - 1); i++ {
		path *= i
		path /= (i - n + 1)
	}
	return path
}

func main() {

}
