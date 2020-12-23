package josefhu

import (
	"fmt"
	"testing"
)

func Test_josephuLink(t *testing.T) {
	type args struct {
		m int
		k int
		n int
	}
	tests := []struct {
		name string
		args args
		want []int
	}{
		{name: "success", args: args{3, 1, 5}},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := josephuLink(tt.args.m, tt.args.k, tt.args.n)
			fmt.Println(got)
		})
	}
}
