package Day4

type Node struct {
	Name     string
	Children []*Node
}

func (n *Node) depthFirstSearch(array []string) []string {
	array = append(array, n.Name)
	for _, child := range n.Children {
		array = child.depthFirstSearch(array)
	}
	return array
}
