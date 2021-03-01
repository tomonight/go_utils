package tree

var queue = [100]int{}
var kmax = 1

func isCompleteTree(root *TreeNode) bool {

	for i := 0; i < 100; i++ {
		queue[i] = -2
	}
	if root == nil {
		return true
	}
	queue[0] = root.Val
	com(root, 0)
	for i := 1; i <= kmax; i++ {
		if (queue[i] == -2 && i != kmax) || (queue[i] == -1 && i != kmax) {
			return false
		}
	}

	return true

}

var a = []int{}

func isValidBST(root *TreeNode) bool {
	a = []int{}
	if root == nil {
		return true
	}
	d(root)
	for i := 0; i < len(a); i++ {
		for j := i + 1; j < len(a); j++ {
			if a[i] < a[j] {
				return false
			}
		}
	}
	return true
}

func d(root *TreeNode) {
	if root == nil {
		return
	}
	d(root.Left)
	a = append(a, root.Val)
	d(root.Right)
}

func com(root *TreeNode, n int) {

	if root == nil {
		return
	}
	if kmax < n {
		kmax = n
	}
	queue[2*n+1] = getV(root.Left)
	queue[2*n+2] = getV(root.Right)

	com(root.Left, 2*n+1)
	com(root.Right, 2*n+2)
}

func getV(root *TreeNode) int {
	if root == nil {
		return -1
	}
	return root.Val
}

func pre(root *TreeNode) []int {
	if root == nil {
		return nil
	}
	re := []int{}
	stack := []*TreeNode{}
	stack = append(stack, root)
	for len(stack) != 0 {
		p := stack[len(stack)-1]
		re = append(re, p.Val)
		stack = stack[:len(stack)-1]
		if p.Right != nil {
			stack = append(stack, p.Right)
		}
		if p.Left != nil {
			stack = append(stack, p.Left)
		}

	}
	return re
}

func mid(root *TreeNode) []int {
	if root == nil {
		return nil
	}
	tmp := root
	re := []int{}
	stack := []*TreeNode{}
	//stack = append(stack,root)
	for len(stack) != 0 || tmp != nil {
		for tmp != nil {
			stack = append(stack, tmp)
			tmp = tmp.Left
		}
		tmp = stack[len(stack)-1]
		re = append(re, tmp.Val)
		stack = stack[:len(stack)-1]
		tmp = tmp.Right
	}
	return re
}
