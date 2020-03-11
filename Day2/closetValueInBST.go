package Day2

import "math"

type BST struct {
	value int

	left  *BST
	right *BST
}

//Average: O(log(n)) time | O(log(n)) space
//Worst: O(n) time | O(n) space

func (tree *BST) FindClosestValue(target int) int {
	return tree.findClosestValue(target, math.MaxInt32)
}

func (tree *BST) findClosestValue(target, closest int) int {
	if absdiff(target, closest) > absdiff(target, tree.value) {
		closest = tree.value
	}
	if target < tree.value && tree.left != nil {
		return tree.left.findClosestValue(target, closest)
	} else if target > tree.value && tree.right != nil {
		return tree.right.findClosestValue(target, closest)
	}
	return closest
}

func absdiff(a, b int) int {
	out := math.Abs(float64(a) - float64(b))
	return int(out)
}
