package closedisland

func closedIsland(grid [][]int) int {
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
			if used[i][j] || grid[i][j] == 1 {
				continue
			}

			if checkcloseLandDFS(grid, used, i, j) {
				count++
			}
		}
	}

	return count
}

func checkcloseLandDFS(grid [][]int, used [][]bool, posX, posY int) bool {
	flag := true
	rowMax, colMax := len(grid), len(grid[0])
	searchQueue := make([]Point, 0)

	searchQueue = append(searchQueue, Point{x: posX, y: posY})
	used[posX][posY] = true

	for len(searchQueue) > 0 {
		curPoint := searchQueue[0]
		searchQueue = searchQueue[1:]

		if curPoint.x == 0 || curPoint.x == rowMax-1 ||
			curPoint.y == 0 || curPoint.y == colMax-1 {
			flag = false
		}

		for _, v := range AroundPoints {
			checkPoint := Point{
				x: curPoint.x + v.x,
				y: curPoint.y + v.y,
			}

			if (checkPoint.x >= 0 && checkPoint.x < rowMax) &&
				(checkPoint.y >= 0 && checkPoint.y < colMax) &&
				grid[checkPoint.x][checkPoint.y] == 0 &&
				!used[checkPoint.x][checkPoint.y] {
				searchQueue = append(searchQueue, checkPoint)
				used[checkPoint.x][checkPoint.y] = true

				if checkPoint.x == 0 || checkPoint.x == rowMax-1 ||
					checkPoint.y == 0 || checkPoint.y == colMax-1 {
					flag = false
				}
			}
		}
	}

	return flag
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
