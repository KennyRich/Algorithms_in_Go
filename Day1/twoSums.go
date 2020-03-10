package Day1
import "sort"

//O(n) time | O(n) space
func TwoNumberSum(array[]int, target int) [] int {
	numToIndex := map[int]bool{}
	for _, num := range array{
		potentialMatch := target - num
		if _, found := numToIndex[potentialMatch]; found{
			return []int {potentialMatch, num}
		}else{
			numToIndex[num] = true
		}
	}
	return []int {}
}

// O(nlog(n)) time | O(1) space
func addTwoNumbers(array[]int, target int) [] int {
	sort.Ints(array)
	left, right := 0, len(array)-1
	for left < right {
		currentSum := array[left] + array[right]
		if currentSum == target{
			return []int {array[left], array[right]}
		}else if currentSum < target{
			left +=1
		}else{
			right -=1
		}
	}
	return []int {}
}
