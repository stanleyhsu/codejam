package leetcode

// Definition for a binary tree node.
type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func appendSlice(s []*TreeNode, n *TreeNode) []*TreeNode {
	if n.Left != nil {
		s = append(s, n.Left)
	}
	if n.Right != nil {
		s = append(s, n.Right)
	}
	return s
}

func isMirror(n1, n2 *TreeNode) bool {
	if n1 == nil && n2 == nil {
		return true
	}
	if n1 == nil || n2 == nil {
		return false
	}
	if n1.Val != n2.Val {
		return false
	}
	if (n1.Left == nil && n2.Right != nil) || (n1.Left != nil && n2.Right == nil) {
		return false
	}
	if (n1.Right == nil && n2.Left != nil) || (n1.Right != nil && n2.Left == nil) {
		return false
	}
	return true
}
func isSymmetric(root *TreeNode) bool {
	if root == nil {
		return true
	}

	nodeSlice := make([]*TreeNode, 0, 8)
	nodeSlice = appendSlice(nodeSlice, root)

	for len(nodeSlice) > 0 {
		n := len(nodeSlice)
		if n%2 != 0 {
			return false
		}
		for i := 0; i < n/2; i++ {
			n1, n2 := nodeSlice[i], nodeSlice[n-1-i]
			if !isMirror(n1, n2) {
				return false
			}

		}
		for i := 0; i < n; i++ {
			nodeSlice = appendSlice(nodeSlice, nodeSlice[i])
		}
		nodeSlice = nodeSlice[n:]
	}

	return true
}

func isMirrorv2(n1, n2 *TreeNode) bool {
	if n1 == nil && n2 == nil {
		return true
	}
	if n1 == nil || n2 == nil {
		return false
	}
	if n1.Val != n2.Val {
		return false
	}
	return isMirrorv2(n1.Left, n2.Right) && isMirrorv2(n1.Right, n2.Left)
}

func isSymmetricv2(root *TreeNode) bool {
	if root == nil {
		return true
	}
	return isMirrorv2(root.Left, root.Right)
}
