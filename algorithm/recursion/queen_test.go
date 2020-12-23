package recursion

import (
	"fmt"
	"testing"
)

func Test_princess(t *testing.T) {
	type args struct {
		q [8][8]int
		n int
	}
	tests := []struct {
		name string
		args args
		want bool
	}{
		{name: "success"},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			q := [8]int{}
			i := 0
			princess(q, 0, &i)
			fmt.Println(i)

			// queen(q, 0)
		})
	}
}
