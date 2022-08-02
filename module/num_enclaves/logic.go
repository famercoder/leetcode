package numenclaves

func numEnclaves(grid [][]int) int {
	if len(grid) == 0 {
		return 0
	}

	row, col := len(grid), len(grid[0])
	used := make([][]bool, row)

	for i := 0; i < row; i++ {
		used[i] = make([]bool, col)
	}

	num := 0

	for i := 0; i < row; i++ {
		for j := 0; j < col; j++ {
			if used[i][j] || grid[i][j] == 0 {
				continue
			}

			num += checkEnclaves(grid, used, i, j)
		}
	}

	return num
}

func checkEnclaves(grid [][]int, used [][]bool, posX, posY int) int {
	row, col := len(grid), len(grid[0])
	searchQueue := make([]Point, 0)

	searchQueue = append(searchQueue, Point{x: posX, y: posY})
	used[posX][posY] = true

	flag := false
	num := 0

	if posX == 0 || posX == row-1 ||
		posY == 0 || posY == col-1 {
		flag = true
	}

	for len(searchQueue) > 0 {
		curPoint := searchQueue[0]
		searchQueue = searchQueue[1:]
		num++

		for _, v := range AroundPoints {
			checkPoint := Point{
				x: curPoint.x + v.x,
				y: curPoint.y + v.y,
			}

			if (checkPoint.x >= 0 && checkPoint.x < row) &&
				(checkPoint.y >= 0 && checkPoint.y < col) &&
				grid[checkPoint.x][checkPoint.y] == 1 &&
				!used[checkPoint.x][checkPoint.y] {
				searchQueue = append(searchQueue, checkPoint)
				used[checkPoint.x][checkPoint.y] = true

				if checkPoint.x == 0 || checkPoint.x == row-1 ||
					checkPoint.y == 0 || checkPoint.y == col-1 {
					flag = true
				}
			}
		}
	}

	if flag {
		return 0
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
