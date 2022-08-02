package permute

//这是一个排列算法，获取nums全排列
func permute(nums []int) [][]int {
	arrs := make([][]int, 0)
	if len(nums) == 0 {
		return arrs
	}

	used := make([]bool, len(nums))
	path := make([]int, 0)

	perm(nums, used, &arrs, &path)
	return arrs
}

func perm(nums []int, used []bool, arrs *[][]int, path *[]int) {
	if len(*path) == len(nums) {
		arr := make([]int, len(nums))
		copy(arr, *path)

		*arrs = append(*arrs, arr)
		return
	}

	for i := 0; i < len(nums); i++ {
		if used[i] {
			continue
		}

		*path = append(*path, nums[i])
		used[i] = true

		perm(nums, used, arrs, path)

		*path = (*path)[:len(*path)-1]
		used[i] = false
	}
}
