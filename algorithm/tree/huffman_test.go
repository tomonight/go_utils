package tree

import (
	"fmt"
	"testing"
)

func Test_huffmanTree(t *testing.T) {
	type args struct {
		data [][]interface{}
	}
	tests := []struct {
		name string
		args args
		want *TreeNode
	}{
		{name: "success", args: args{
			data: [][]interface{}{
				[]interface{}{1, byte('l')},
				[]interface{}{100, byte('i')},
				[]interface{}{20, byte(' ')},
				[]interface{}{4, byte('e')},
				[]interface{}{6, byte('i')},
				[]interface{}{10, byte('j')},
				[]interface{}{20, byte('a')},
				[]interface{}{2, byte('d')}},
		},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := huffmanTree(tt.args.data)

			makeHuffmanMap(got, "")
			getMap()
			fmt.Println(got)
		})
	}
}
