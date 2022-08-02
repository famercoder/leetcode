package combine

//这是一个组合算法，获取1~n之间选k个数的组合
func combine(n int, k int) [][]int {
	arrs := make([][]int, 0)
	queue := make([]int, 0)

	if k <= 0 || n < k {
		return arrs
	}

	CombineDFS(n, k, 1, &arrs, &queue)
	return arrs
}

func CombineDFS(n, k, begin int, arrs *[][]int, queue *[]int) {
	if len(*queue) == k {
		arr := make([]int, len(*queue))
		for i := 0; i < len(*queue); i++ {
			arr[i] = (*queue)[i]
		}

		*arrs = append(*arrs, arr)
		return
	}

	for i := begin; i <= n; i++ {
		*queue = append(*queue, i)
		CombineDFS(n, k, i+1, arrs, queue)
		*queue = (*queue)[:len(*queue)-1]
	}
}

//组合算法,获取srcArr中k个数的组合
func combineArrs(srcArr []int, k int) [][]int {
	arrs := make([][]int, 0)
	queue := make([]int, 0)

	if len(srcArr) == 0 || len(srcArr) < k {
		return arrs
	}

	combineArrsDFS(srcArr, k, 0, &arrs, &queue)
	return arrs
}

func combineArrsDFS(srcArr []int, k, begin int, arrs *[][]int, queue *[]int) {
	if len(*queue) == k {
		arr := make([]int, len(*queue))
		copy(arr, *queue)

		*arrs = append(*arrs, arr)
		return
	}

	for i := begin; i < len(srcArr); i++ {
		*queue = append(*queue, srcArr[i])
		combineArrsDFS(srcArr, k, i+1, arrs, queue)
		*queue = (*queue)[:len(*queue)-1]
	}
}
