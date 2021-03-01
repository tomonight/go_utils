package tree

import (
	"fmt"
	"testing"
)

func Test_maxPathSum(t *testing.T) {
	root := &TreeNode{Val: -10}
	left := &TreeNode{Val: 9}
	right := &TreeNode{Val: 20}
	rightL := &TreeNode{Val: 15}
	rightR := &TreeNode{Val: 7}
	right.Left = rightL
	right.Right = rightR
	root.Right = right
	root.Left = left
	type args struct {
		root *TreeNode
	}

	tests := []struct {
		name string
		args args
		want int
	}{
		{name: "success", args: args{root: root}, want: -1},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			//if got := maxPathSum(tt.args.root); got != tt.want {
			//	t.Errorf("maxPathSum() = %v, want %v", got, tt.want)
			//}
			fmt.Println(twoSum([]int{3, 2, 4}, 6))
		})
	}
}

func Test_isCompleteTree(t *testing.T) {
	root := &TreeNode{Val: 2}
	left := &TreeNode{Val: 3}
	right := &TreeNode{Val: 1}
	leftL := &TreeNode{Val: 3}
	leftR := &TreeNode{Val: 6}
	rightR := &TreeNode{Val: 7}
	left.Left = leftL
	left.Right = leftR
	right.Right = rightR
	root.Right = left
	root.Left = right
	type args struct {
		root *TreeNode
	}

	tests := []struct {
		name string
		args args
		want int
	}{
		{name: "success", args: args{root: root}, want: -1},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			a := isValidBST(root)
			fmt.Println(a)
		})
	}
}
