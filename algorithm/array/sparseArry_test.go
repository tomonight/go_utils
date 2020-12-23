package algorithm

import (
	"fmt"
	"testing"
)

func TestSparseArryInt(t *testing.T) {
	type args struct {
		arry       [][]int
		spaseValue int
	}
	tests := []struct {
		name string
		args args
		want [][3]int
	}{
		{name: "success", args: args{arry: [][]int{{0, 0, 0, 0}, {0, 1, 0, 0}, {0, 2, 0, 0}}, spaseValue: 0}},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := SparseArryInt(tt.args.arry, tt.args.spaseValue)

			for _, raw := range got {
				raw.Print()
			}

			got1 := RecoverArryInt(got, tt.args.spaseValue)

			for _, raw := range got1 {
				for _, va := range raw {
					fmt.Printf("%d  ", va)
				}
				fmt.Println()
			}
		})
	}
}
