package maxdistance

func maxDistance(grid [][]int) int {
	if len(grid) == 0 {
		return -1
	}

	row, col := len(grid), len(grid[0])
	landQueue := make([]Point, 0)

	for i := 0; i < row; i++ {
		for j := 0; j < col; j++ {
			if grid[i][j] == 1 {
				landQueue = append(landQueue, Point{x: i, y: j})
			}
		}
	}

	hasSea := false
	point := Point{x: -1, y: -1}

	for len(landQueue) > 0 {
		point = landQueue[0]
		landQueue = landQueue[1:]

		for _, v := range AroundPoints {
			checkPoint := Point{
				x: point.x + v.x,
				y: point.y + v.y,
			}

			if (checkPoint.x >= 0 && checkPoint.x < row) &&
				(checkPoint.y >= 0 && checkPoint.y < col) &&
				grid[checkPoint.x][checkPoint.y] == 0 {
				hasSea = true
				landQueue = append(landQueue, checkPoint)
				grid[checkPoint.x][checkPoint.y] = grid[point.x][point.y] + 1
			}
		}
	}

	if point.x == -1 || !hasSea {
		return -1
	}

	return grid[point.x][point.y] - 1
}

type Point struct {
	x int
	y int
}

var AroundPoints = []Point{
	{-1, 0}, //up
	{1, 0},  //down
	{0, -1}, //left
	{0, 1},  //right
}
