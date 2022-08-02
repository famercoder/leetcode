package maxareaofisland

func maxAreaOfIsland(grid [][]int) int {
	rowMax := len(grid)
	colMax := len(grid[0])

	gridFlag := make([][]bool, rowMax)
	for i := 0; i < rowMax; i++ {
		gridFlag[i] = make([]bool, colMax)
	}

	maxLand := 0

	for i := 0; i < rowMax; i++ {
		for j := 0; j < colMax; j++ {
			if !gridFlag[i][j] && grid[i][j] == 1 {
				land := searchIsland(grid, gridFlag, i, j)
				if land > maxLand {
					maxLand = land
				}
			}
		}
	}

	return maxLand
}

func searchIsland(grid [][]int, gridFlag [][]bool, posX, posY int) int {
	rowMax := len(grid)
	colMax := len(grid[0])

	fillPoints := make([]Point, 0)
	startPoint := Point{x: posX, y: posY}
	fillPoints = append(fillPoints, startPoint)

	searchedPoints := map[Point]struct{}{}
	searchedPoints[startPoint] = struct{}{}

	for len(fillPoints) > 0 {
		point := fillPoints[0]
		fillPoints = fillPoints[1:]
		gridFlag[point.x][point.y] = true

		for _, v := range AroundPoints {
			checkPoint := Point{
				x: point.x + v.x,
				y: point.y + v.y,
			}

			_, find := searchedPoints[checkPoint]
			if (checkPoint.x >= 0 && checkPoint.x < rowMax) &&
				(checkPoint.y >= 0 && checkPoint.y < colMax) &&
				(grid[checkPoint.x][checkPoint.y] == 1 && !find) {
				fillPoints = append(fillPoints, checkPoint)
				searchedPoints[checkPoint] = struct{}{}
			}
		}
	}

	return len(searchedPoints)
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
