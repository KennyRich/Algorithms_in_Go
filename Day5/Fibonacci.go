package Day5

func getNthFib(n int) int {
	lastTwo := []int{0, 1}
	counter := 3

	for counter <= n {
		nextFib := lastTwo[0] + lastTwo[1]
		lastTwo = []int{lastTwo[1], nextFib}
		counter += 1
	}
	if n > 1 {
		return lastTwo[1]
	} else {
		return lastTwo[0]
	}
}
