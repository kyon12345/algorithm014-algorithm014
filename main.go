package main

import "fmt"

var count int

func totalNQueens(n int) int {
	count = 0

	dfs(n, 0, 0, 0, 0)

	return count
}

func dfs(n, row, col, right, left int) {
	if n == row {
		count++
		return
	}

	availablePositions := ^(col | right | left) & (1<<n - 1)

	for availablePositions > 0 {
		pick := availablePositions & -availablePositions
		fmt.Printf("%08b\n", pick)
		dfs(n, row+1, col|pick, (right|pick)>>1, (left|pick)<<1)
		//revert
		availablePositions &= availablePositions - 1
	}
}

func main() {
	totalNQueens(4)
}
