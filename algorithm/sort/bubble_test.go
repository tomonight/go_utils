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
	oo := []int{803, 164, 576, 978, 20, 617, 986, 365, 675, 537, 824, 535, 580, 428, 971, 947, 299, 19, 139, 534, 412, 69, 759, 235, 285, 390, 295, 307, 190, 54, 366}
	tests := []struct {
		name string
		args args
	}{
		{name: "success", args: args{data: oo}},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			d := tt.args.data
			a := []int{}
			x := 8000
			for i := 0; i < x; i++ {
				a = append(a, rand.Intn(x))
			}
			b := time.Now().Unix()
			// shellAsc(tt.args.data)
			QuickAsc(d)
			//fmt.Println(kk)
			c := time.Now().Unix()
			// fmt.Println(a)
			// fmt.Println(tt.args.data)
			fmt.Println(c - b)
		})
	}
}
