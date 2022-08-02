package floodfill

// input:image = [[1,1,1],[1,1,0],[1,0,1]],sr = 1, sc = 1, newColor = 2
// output:[[2,2,2],[2,2,0],[2,0,1]]
func floodFill(image [][]int, sr int, sc int, color int) [][]int {
	rowMax := len(image)
	colMax := len(image[0])

	fillPoints := make([]Point, 0)
	fillPoints = append(fillPoints, Point{sr, sc})

	searchedPoints := map[Point]struct{}{}
	startColor := image[sr][sc]

	for len(fillPoints) > 0 {
		point := fillPoints[0]
		fillPoints = fillPoints[1:]
		image[point.x][point.y] = color

		for _, v := range AroundPoints {
			checkPoint := Point{
				x: point.x + v.x,
				y: point.y + v.y,
			}

			_, find := searchedPoints[checkPoint]
			if (checkPoint.x >= 0 && checkPoint.x < rowMax) &&
				(checkPoint.y >= 0 && checkPoint.y < colMax) &&
				(image[checkPoint.x][checkPoint.y] == startColor && !find) {
				fillPoints = append(fillPoints, checkPoint)
				searchedPoints[checkPoint] = struct{}{}
			}
		}
	}

	return image
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
