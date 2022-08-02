package numenclaves

import (
	"fmt"
	"testing"
)

func TestNumEnclaves(t *testing.T) {
	grig := [][]int{
		{0, 0, 0, 0},
		{1, 0, 1, 0},
		{0, 1, 1, 0},
		{0, 0, 0, 0},
	}
	num := numEnclaves(grig)
	fmt.Println(num)
}
