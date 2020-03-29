package Day3

type BinaryTree struct {
	value int
	left  *BinaryTree
	right *BinaryTree
}

func BranchSums(root *BinaryTree) []int {
	sums := []int{}
	calculateBranchSums(root, 0, &sums)
	return sums
}

func calculateBranchSums(node *BinaryTree, runningSum int, sums *[]int) {
	if node == nil {
		return
	}
	runningSum += node.value
	if node.left == nil && node.right == nil {
		*sums = append(*sums, runningSum)
		return
	}
	calculateBranchSums(node.left, runningSum, sums)
	calculateBranchSums(node.right, runningSum, sums)
}
