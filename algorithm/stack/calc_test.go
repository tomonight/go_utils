package stack

import (
	"fmt"
	"testing"
)

func Test_toMidExpr(t *testing.T) {
	type args struct {
		expr string
	}
	tests := []struct {
		name string
		args args
		want []map[int]string
	}{
		{name: "success", args: args{expr: "1*(2+3)*4+2*10/(5*2)+100"}},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := toMidExpr(tt.args.expr)
			fmt.Println(got)
			mm := toSuffix(got)
			fmt.Println(mm)

			kk := calcStack(mm)
			fmt.Println(kk)
		})
	}
}
