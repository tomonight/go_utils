package basic

import "testing"

func Test_bac(t *testing.T) {
	type args struct {
		num int
		a   string
		b   string
		c   string
	}
	tests := []struct {
		name string
		args args
	}{
		{name: "success", args: args{}},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			letterCombinations("234")
		})
	}
}

func getByte(a byte) byte {
	return 97 + (a-48)*3
}
