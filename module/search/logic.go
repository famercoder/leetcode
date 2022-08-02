package search

func search(nums []int, target int) int {
	length := len(nums)
	low, high := 0, length-1

	for low <= high {
		mid := (low + high) / 2
		if nums[mid] > target {
			high = mid - 1
		} else if nums[mid] < target {
			low = mid + 1
		} else {
			return mid
		}
	}

	return -1
}
