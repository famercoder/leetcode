package numsislands

import (
	"fmt"
	"testing"
)

func TestNumIslands(t *testing.T) {
	grid := [][]byte{
		{1, 1, 1, 1, 0},
		{1, 1, 0, 1, 0},
		{1, 1, 0, 0, 0},
		{0, 0, 0, 0, 0},
	}

	num := numIslands(grid)
	fmt.Println("result:", num)
}
