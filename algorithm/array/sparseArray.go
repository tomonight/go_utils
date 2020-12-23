package algorithm

import "fmt"

type Sparse struct {
	row int
	col int
	val int
}

func (s *Sparse) Print() {
	fmt.Println(fmt.Sprintf("%d  %d   %d", s.col, s.row, s.val))

}
func SparseArryInt(arry [][]int, spaseValue int) []Sparse {

	sparseArry := []Sparse{}
	count := 0
	sparseArry = append(sparseArry, Sparse{len(arry), len(arry[0]), count})
	for row, c1 := range arry {
		for col, value := range c1 {
			if value != spaseValue {
				count++
				val := Sparse{row, col, value}
				sparseArry = append(sparseArry, val)
			}
		}
	}
	sparseArry[0].val = count
	return sparseArry
}

func RecoverArryInt(arry []Sparse, sparseValue int) [][]int {

	var recoverArry [][]int

	for i := 1; i < len(arry); i++ {
		rowValue := arry[i]
		if rowValue.val != sparseValue {
			recoverArry[rowValue.col][rowValue.row] = rowValue.val
		}
	}
	return recoverArry
}
