package Day9

// Average time complexity: O(n) | Space Complexity O(1)
func bubbleSort(array []int) []int {
	isSorted := false
	counter := 0
	for !isSorted {
		isSorted = true
		for i := 0; i < len(array)-1-counter; i++ {
			if array[i] > array[i+1] {
				array[i], array[i+1] = array[i+1], array[i]
				isSorted = false
			}
		}
		counter++
	}
	return array
}
