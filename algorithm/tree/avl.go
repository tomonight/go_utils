package tree

var tmpMax = 0

func newAVLTree(data []int) *TreeNode {

	root := getBST(data)

	return avlTree(root)
}

func avlTree(root *TreeNode) *TreeNode {
	tmpRoot := root
	i, j := treeDeep(root.Left, 0), treeDeep(root.Right, 0)
	for j-i > 1 {
		tmpRoot = cycleLeft(tmpRoot)
		i, j = treeDeepth(tmpRoot.Left), treeDeepth(tmpRoot.Right)
	}
	for i-j > 1 {
		tmpRoot = cycleRight(tmpRoot)
		i, j = treeDeepth(tmpRoot.Left), treeDeepth(tmpRoot.Right)
	}
	if tmpRoot.Left != nil {
		il, jl := treeDeep(tmpRoot.Left.Left, 0), treeDeep(tmpRoot.Left.Right, 0)
		if jl-il > 1 || il-jl > 1 {
			tmpRoot.Left = avlTree(tmpRoot.Left)
		}
	}
	if tmpRoot.Right != nil {
		il, jl := treeDeep(tmpRoot.Right.Left, 0), treeDeep(tmpRoot.Right.Right, 0)
		if jl-il > 1 || il-jl > 1 {
			tmpRoot.Left = avlTree(tmpRoot.Left)
		}
	}

	return tmpRoot
}

func cycleLeft(tmpRoot *TreeNode) *TreeNode {
	if m, n := treeDeepth(tmpRoot.Right.Left), treeDeepth(tmpRoot.Right.Right); m > n && n > 0 {
		tmpRoot.Right = cycleRight(tmpRoot.Right)
	}
	t := tmpRoot
	tmpNode := &TreeNode{Val: tmpRoot.Val}
	tmpRoot = t.Right
	tmpNode.Right = tmpRoot.Left
	tmpNode.Left = t.Left
	tmpRoot.Left = tmpNode
	return tmpRoot
	//i,j = treeDeep(tmpRoot.Left,0),treeDeep(tmpRoot.Right,0)
}

func cycleRight(tmpRoot *TreeNode) *TreeNode {
	if m, n := treeDeepth(tmpRoot.Left.Right), treeDeepth(tmpRoot.Left.Left); m > n && n > 0 {
		tmpRoot.Left = cycleLeft(tmpRoot.Left)
	}
	t := tmpRoot
	tmpNode := &TreeNode{Val: tmpRoot.Val}
	tmpRoot = t.Left
	tmpNode.Right = tmpRoot.Left
	tmpNode.Right = t.Right
	tmpRoot.Right = tmpNode
	return tmpRoot
}

func treeDeepth(root *TreeNode) int {
	return treeDeep(root, 0)
}
func treeDeep(root *TreeNode, deep int) int {
	if root == nil {
		return deep
	}
	if root.Left == nil && root.Right == nil {
		//tmpMax = umax(tmpMax,deep + 1)
		return deep + 1
	}

	offerl := treeDeep(root.Left, deep)
	offerr := treeDeep(root.Right, deep)

	return umax(offerl, offerr) + 1
}

func umax(i, j int) int {
	if i > j {
		return i
	}
	return j
}
