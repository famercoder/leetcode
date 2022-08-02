package pacificatlantic

func pacificAtlantic(heights [][]int) [][]int {
	if len(heights) == 0 {
		return nil
	}

	row, col := len(heights), len(heights[0])
	results := make([][]int, 0)

	var bfsFn func(point Point, row, col int, results *[][]int, heights [][]int)
	bfsFn = func(point Point, row, col int, results *[][]int, heights [][]int) {
		searchQueue := make([]Point, 0)
		searchQueue = append(searchQueue, point)

		used := make([][]bool, row)
		for i := 0; i < row; i++ {
			used[i] = make([]bool, col)
		}

		used[point.x][point.y] = true
		pacificOcean, atlanticOcean := false, false

		for len(searchQueue) > 0 {
			curPoint := searchQueue[0]
			searchQueue = searchQueue[1:]

			if curPoint.x == 0 || curPoint.y == 0 {
				pacificOcean = true
			}

			if curPoint.x == row-1 || curPoint.y == col-1 {
				atlanticOcean = true
			}

			if pacificOcean && atlanticOcean {
				*results = append(*results, []int{point.x, point.y})
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
					heights[checkPoint.x][checkPoint.y] <= heights[curPoint.x][curPoint.y] {
					searchQueue = append(searchQueue, checkPoint)
					used[checkPoint.x][checkPoint.y] = true
				}
			}
		}
	}

	for i := 0; i < row; i++ {
		for j := 0; j < col; j++ {
			bfsFn(Point{x: i, y: j}, row, col, &results, heights)
		}
	}

	return results
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
