package nearestexit

import (
	"fmt"
	"testing"
)

func TestNearestExit(t *testing.T) {
	maze := [][]byte{
		{'+', '+', '+'},
		{'.', '.', '.'},
		{'+', '+', '+'},
	}

	value := nearestExit(maze, []int{1, 0})
	fmt.Println("Test result:", value)
}
