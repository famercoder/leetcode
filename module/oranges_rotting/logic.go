package orangesrotting

func orangesRotting(grid [][]int) int {
	searchQueue := make([]Point, 0)
	freshQueue := make(map[Point]struct{})

	rowMax := len(grid)
	colMax := len(grid[0])
	num := 0

	for i := 0; i < rowMax; i++ {
		for j := 0; j < colMax; j++ {
			if grid[i][j] == 2 {
				searchQueue = append(searchQueue, Point{x: i, y: j})
			} else if grid[i][j] == 1 {
				freshQueue[Point{x: i, y: j}] = struct{}{}
			}
		}
	}

	if len(freshQueue) == 0 {
		return 0
	}

	for len(searchQueue) > 0 {
		tempQueue := make([]Point, len(searchQueue))
		copy(tempQueue, searchQueue)
		searchQueue = make([]Point, 0)

		for i := 0; i < len(tempQueue); i++ {
			curPoint := tempQueue[i]
			for _, v := range AroundPoints {
				checkPoint := Point{
					x: curPoint.x + v.x,
					y: curPoint.y + v.y,
				}

				if (checkPoint.x >= 0 && checkPoint.x < rowMax) &&
					(checkPoint.y >= 0 && checkPoint.y < colMax) &&
					grid[checkPoint.x][checkPoint.y] == 1 {
					searchQueue = append(searchQueue, checkPoint)
					delete(freshQueue, checkPoint)
					grid[checkPoint.x][checkPoint.y] = 2
				}
			}
		}

		if len(searchQueue) > 0 {
			num++
		}
	}

	if len(freshQueue) > 0 {
		return -1
	}

	return num
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
