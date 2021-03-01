package search

import (
	"fmt"
	"testing"
)

func Test_binarySearch(t *testing.T) {
	type args struct {
		arr   []int
		val   int
		left  int
		right int
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{name: "success", args: args{arr: []int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10, 20, 30, 40, 50, 60}, val: 60}},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := noRecursion(tt.args.arr, tt.args.val)
			fmt.Println(got)
		})
	}
}
