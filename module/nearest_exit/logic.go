package nearestexit

//maze = [["+","+",".","+"],[".",".",".","+"],["+","+","+","."]]
//entrance = [1,2]
//out:1
func nearestExit(maze [][]byte, entrance []int) int {
	if len(maze) == 0 || len(entrance) != 2 {
		return -1
	}

	row, col := len(maze), len(maze[0])
	used := make([][]bool, row)
	grid := make([][]int, row)

	for i := 0; i < row; i++ {
		used[i] = make([]bool, col)
		grid[i] = make([]int, col)

		for j := 0; j < col; j++ {
			if maze[i][j] == '+' {
				grid[i][j] = 1
			}
		}
	}

	exitPoint := Point{x: -1, y: -1}
	entrancePoint := Point{x: entrance[0], y: entrance[1]}

	bfsFn := func() {
		searchQueue := make([]Point, 0)
		searchQueue = append(searchQueue, entrancePoint)
		used[entrancePoint.x][entrancePoint.y] = true

		for len(searchQueue) > 0 {
			curPoint := searchQueue[0]
			searchQueue = searchQueue[1:]

			if curPoint != entrancePoint &&
				(curPoint.x == 0 || curPoint.y == 0 ||
					curPoint.x == row-1 || curPoint.y == col-1) {
				exitPoint = curPoint
				break
			}

			for _, v := range AroundPoints {
				checkPoint := Point{
					x: curPoint.x + v.x,
					y: curPoint.y + v.y,
				}

				if (checkPoint.x >= 0 && checkPoint.x < row) &&
					(checkPoint.y >= 0 && checkPoint.y < col) &&
					!used[checkPoint.x][checkPoint.y] &&
					grid[checkPoint.x][checkPoint.y] == 0 {
					searchQueue = append(searchQueue, checkPoint)
					used[checkPoint.x][checkPoint.y] = true
					grid[checkPoint.x][checkPoint.y] = grid[curPoint.x][curPoint.y] + 1
				}
			}
		}
	}

	bfsFn()

	if exitPoint.x == -1 {
		return -1
	}

	return grid[exitPoint.x][exitPoint.y]
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
