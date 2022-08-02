package floodfill

import (
	"fmt"
	"testing"
)

func TestFloodFill(t *testing.T) {
	image := [][]int{
		{1, 1, 1},
		{1, 1, 0},
		{1, 0, 1},
	}

	results := floodFill(image, 1, 1, 2)
	fmt.Println("result:", results)
}
