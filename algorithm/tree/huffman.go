package tree

import "fmt"

type HuffmanTree struct {
	Val   int
	b     byte
	Left  *HuffmanTree
	Right *HuffmanTree
}

var (
	huffmanMap map[byte]string = make(map[byte]string)
)

func makeHuffmanMap(tree *HuffmanTree, code string) {
	if tree == nil {
		return
	}
	if tree.Left == nil && tree.Right == nil {
		huffmanMap[tree.b] = code
	}
	makeHuffmanMap(tree.Left, code+"0")
	makeHuffmanMap(tree.Right, code+"1")
}

func getMap() {
	fmt.Println(huffmanMap)
}

func huffmanTree(data [][]interface{}) *HuffmanTree {
	trees := []*HuffmanTree{}
	for _, v := range data {
		tree := &HuffmanTree{Val: v[0].(int), b: v[1].(byte)}
		trees = append(trees, tree)
	}
	var r *HuffmanTree
	QuickAsc(trees)
	tmp := trees
	for len(tmp) > 1 {
		a := tmp[0]
		b := tmp[1]
		r = &HuffmanTree{Val: a.Val + b.Val, Left: a, Right: b}
		tmp = append(tmp[2:], r)
		QuickAsc(tmp)
	}
	return r
}

func sortTree(trees []*HuffmanTree) []*HuffmanTree {
	if len(trees) == 0 {
		return trees
	}
	n := []*HuffmanTree{}
	for i := 0; i < len(trees); i++ {
		min := trees[i]
		for j := i + 1; j < len(trees); j++ {
			if min.Val > trees[j].Val {
				min = trees[j]
			}
		}
		n = append(n, min)
	}
	return n
}

func QuickAsc(arr []*HuffmanTree) {
	quickAsc(arr, 0, len(arr)-1)
}

//
func quickAsc(arr []*HuffmanTree, left, right int) {
	if left == right || right == 0 {
		return
	}
	l, r := left, right
	if l < r {
		k := partition(arr, l, r)
		quickAsc(arr, left, k-1)
		quickAsc(arr, k+1, right)
	}
}

func partition(arr []*HuffmanTree, left, right int) int {
	index := left + 1
	for i := left; i <= right; i++ {
		if arr[i].Val < arr[left].Val {
			arr[i], arr[index] = arr[index], arr[i]
			index++
		}
	}
	arr[index-1], arr[left] = arr[left], arr[index-1]
	return index - 1
}
