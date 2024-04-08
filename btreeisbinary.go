package piscine

func BTreeIsBinary(root *TreeNode) bool {
	return IsBHelper(root, nil, nil)
}

func IsBHelper(node *TreeNode, min, max *TreeNode) bool {
	if node == nil {
		return true
	}
	if (min != nil && node.Data <= min.Data) || (max != nil && node.Data >= max.Data) {
		return false
	}
	return IsBHelper(node.Left, min, node) && IsBHelper(node.Right, node, max)
}
