package numsislands

func numIslands(grid [][]byte) int {
	if len(grid) == 0 {
		return 0
	}

	row, col := len(grid), len(grid[0])
	used := make([][]bool, row)

	for i := 0; i < row; i++ {
		used[i] = make([]bool, col)
	}

	count := 0
	for i := 0; i < row; i++ {
		for j := 0; j < col; j++ {
			if used[i][j] || grid[i][j] == 0 {
				continue
			}

			if isLandsDFS(grid, used, i, j) {
				count++
			}
		}
	}

	return count
}

func isLandsDFS(grid [][]byte, used [][]bool, posX, posY int) bool {
	rowMax, colMax := len(grid), len(grid[0])
	searchQueue := make([]Point, 0)

	searchQueue = append(searchQueue, Point{x: posX, y: posY})
	used[posX][posY] = true

	for len(searchQueue) > 0 {
		curPoint := searchQueue[0]
		searchQueue = searchQueue[1:]

		for _, v := range AroundPoints {
			checkPoint := Point{
				x: curPoint.x + v.x,
				y: curPoint.y + v.y,
			}

			if (checkPoint.x >= 0 && checkPoint.x < rowMax) &&
				(checkPoint.y >= 0 && checkPoint.y < colMax) &&
				grid[checkPoint.x][checkPoint.y] == 1 &&
				!used[checkPoint.x][checkPoint.y] {
				searchQueue = append(searchQueue, checkPoint)
				used[checkPoint.x][checkPoint.y] = true
			}
		}
	}

	return true
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
