package piscine

func BTreeApplyByLevel(root *TreeNode, f func(...interface{}) (int, error)) {
	if root == nil {
		return
	}
	MyQueue := []*TreeNode{}
	MyQueue = append(MyQueue, root)
	for len(MyQueue) != 0 {
		current := MyQueue[0]
		MyQueue = pop(MyQueue)
		f(current.Data)
		if current.Left != nil {
			MyQueue = append(MyQueue, current.Left)
		}
		if current.Right != nil {
			MyQueue = append(MyQueue, current.Right)
		}
	}
}

func pop(Btree []*TreeNode) []*TreeNode {
	return Btree[1:]
}
