package logic

//in:rooms = [[1,3],[3,0,1],[2],[0]]
//out:false
func canVisitAllRooms(rooms [][]int) bool {
	if len(rooms) == 0 {
		return false
	}

	row := len(rooms)
	opened := make([]bool, row)

	bfsFn := func() {
		searchQueue := make([]int, 0)
		searchQueue = append(searchQueue, rooms[0]...)
		opened[0] = true

		for len(searchQueue) > 0 {
			roomId := searchQueue[0]
			searchQueue = searchQueue[1:]
			opened[roomId] = true

			for i := 0; i < len(rooms[roomId]); i++ {
				if opened[rooms[roomId][i]] {
					continue
				}

				find := false
				for _, v := range searchQueue {
					if v == rooms[roomId][i] {
						find = true
						break
					}
				}

				if !find {
					searchQueue = append(searchQueue, rooms[roomId][i])
				}
			}
		}
	}

	bfsFn()

	suc := true
	for _, open := range opened {
		if !open {
			suc = false
			break
		}
	}

	return suc
}
