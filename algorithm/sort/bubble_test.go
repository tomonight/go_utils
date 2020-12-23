package sort

import (
	"fmt"
	"math/rand"
	"testing"
	"time"
)

func Test_bubbleAsc(t *testing.T) {
	type args struct {
		data []int
	}
	tests := []struct {
		name string
		args args
	}{
		{name: "success", args: args{data: []int{8, 9, 1, 7, 2, 3, 5, 4, 6, 0}}},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			a := []int{}
			for i := 0; i < 8000000; i++ {
				a = append(a, rand.Intn(8000000))
			}
			b := time.Now().Unix()
			// shellAsc(tt.args.data)
			insertAsc(a)
			c := time.Now().Unix()
			// fmt.Println(a)
			// fmt.Println(tt.args.data)
			fmt.Println(c - b)
		})
	}
}
