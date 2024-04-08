// piscine/btreedeletenode.go
package piscine

// type TreeNode struct {
// 	Data                string
// 	Left, Right, Parent *TreeNode
// }

// // BTreeMin finds the node with the minimum value in a binary search tree
// func BTreeMin(root *TreeNode) *TreeNode {
// 	if root == nil {
// 		return nil
// 	}
// 	for root.Left != nil {
// 		root = root.Left
// 	}
// 	return root
// }

// BTreeDeleteNode deletes a node with a specific value from a binary search tree
func BTreeDeleteNode(root, node *TreeNode) *TreeNode {
	if node != nil {
		if node.Data < root.Data {
			root.Left = BTreeDeleteNode(root.Left, node)
		} else if node.Data > root.Data {
			root.Right = BTreeDeleteNode(root.Right, node)
		} else {
			if root.Left == nil {
				tmp := root.Right
				root = nil
				return tmp
			} else if root.Right == nil {
				tmp := root.Left
				root = nil
				return tmp
			}
			tmp := BTreeMin(root.Right)

			root.Data = tmp.Data
			root.Right = BTreeDeleteNode(root.Right, tmp)
		}
	}

	return root
}

// BTreeSearchItem searches for a specific value in a binary search tree
// func BTreeSearchItem(root *TreeNode, elem string) *TreeNode {
// 	if root == nil {
// 		return nil
// 	}

// 	if elem > root.Data {
// 		return BTreeSearchItem(root.Right, elem)
// 	} else if elem < root.Data {
// 		return BTreeSearchItem(root.Left, elem)
// 	} else if elem == root.Data {
// 		return root
// 	} else {
// 		return root
// 	}
// }

// BTreeIsBinary checks if a given binary tree is a binary search tree
// func BTreeIsBinary(root *TreeNode) bool {
// 	// Implementation of BTreeIsBinary
// 	return true
// }

// BTreeInsertData inserts a new node with the specified data into a binary search tree
// func BTreeInsertData(root *TreeNode, data string) *TreeNode {
// 	// Implementation of BTreeInsertData
// 	return root
// }
