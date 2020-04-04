package Day8

//Time complexity O(n) | Space Complexity O(1)

func binarySearch(array []int, target int) int {
	left := 0
	right := len(array) - 1

	for left <= right {
		mid := (left + right) / 2
		if target == array[mid] {
			return mid
		} else {
			if target < array[mid] {
				right = mid - 1
			} else {
				left = mid + 1
			}
		}
	}
	return -1
}
