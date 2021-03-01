package tree

func getBST(data []int) *TreeNode {
	if len(data) == 0 {
		return nil
	}
	root := &TreeNode{Val: data[0]}
	for i := 1; i < len(data); i++ {
		addNode(root, data[i])
	}
	return root
}

func (b *TreeNode) addBST(val int) {
	tmp := b

	if tmp.Val >= val {
		if tmp.Left == nil {
			tmp.Left = &TreeNode{Val: val}
			return
		}
		tmp.Left.addBST(val)
	} else {
		if tmp.Right == nil {
			tmp.Right = &TreeNode{Val: val}
			return
		}
		tmp.Right.addBST(val)
	}
}

func (b *TreeNode) deleteBST(val int) {

}

func deleteBST(b, parent *TreeNode, val int) {
	var tmpVal *TreeNode
	if b.Val == val {
		if b.Right == nil && b.Left == nil {
			tmpVal = nil
			goto flag
		}
		if b.Left == nil {
			tmpVal = b.Right
			goto flag
		}
		if b.Right == nil {
			tmpVal = b.Left
		}
	flag:
		if b.Val < parent.Val {
			parent.Left = tmpVal
			return
		}
		if parent == nil {
			b = nil
			return
		}
		parent.Right = tmpVal
		return
	}
	if b.Val > val {
		deleteBST(b.Left, b, val)
	} else {
		deleteBST(b.Right, b, val)
	}

}

func addNode(root *TreeNode, val int) {

	if val < root.Val {
		if root.Left != nil {
			addNode(root.Left, val)
			return
		}
		root.Left = &TreeNode{Val: val}
	} else {
		if root.Right != nil {
			addNode(root.Right, val)
			return
		}
		root.Right = &TreeNode{Val: val}
	}
}
