package logic

import (
	"fmt"
	"testing"
)

func TestCanVisitAllRooms(t *testing.T) {
	rooms := [][]int{
		{1, 3},
		{3, 0, 1},
		{2},
		{0},
	}
	flag := canVisitAllRooms(rooms)
	fmt.Println("test result:", flag)
}
