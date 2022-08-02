package updatematrix

func updateMatrix(mat [][]int) [][]int {
	rowMax := len(mat)
	colMax := len(mat[0])

	searchQueue := make([]Point, 0)

	for i := 0; i < rowMax; i++ {
		for j := 0; j < colMax; j++ {
			if mat[i][j] == 0 {
				searchQueue = append(searchQueue, Point{x: i, y: j})
			} else {
				mat[i][j] = -1
			}
		}
	}

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
				mat[checkPoint.x][checkPoint.y] == -1 {
				searchQueue = append(searchQueue, checkPoint)
				mat[checkPoint.x][checkPoint.y] = mat[curPoint.x][curPoint.y] + 1
			}
		}
	}

	return mat
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
