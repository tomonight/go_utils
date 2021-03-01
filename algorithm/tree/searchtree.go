package tree

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

var mmax = 1 << 32
var max = 0

func maxPathSum(root *TreeNode) int {
	max = 0 - mmax
	return max
}

func twoSum(nums []int, target int) []int {

	m := make(map[int]int)
	for i := 0; i < len(nums); i++ {
		m[nums[i]] = i

		if _, ok := m[target-nums[i]]; ok {
			return []int{m[target-nums[i]], i}
		}
	}
	return nil
}

func m(root *TreeNode) int {

	if root == nil {
		return 0
	}
	l := m(root.Left)
	r := m(root.Right)

	lo, ro := 0, 0
	if l > 0 {
		lo = l
	}
	if r > 0 {
		ro = r
	}
	offer := root.Val
	if offer+lo+ro > offer {
		offer = root.Val + lo + ro
	}
	if offer > max {
		max = offer
	}
	return offer
}
