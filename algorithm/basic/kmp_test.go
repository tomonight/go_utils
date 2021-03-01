package basic

import "testing"

func Test_kmp(t *testing.T) {
	type args struct {
		s1 string
		s2 string
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{name: "success", args: args{s1: "BBC ABCDAB ABCDABCDABDE", s2: "ABCDABD"}},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := kmp(tt.args.s1, tt.args.s2); got != tt.want {
				t.Errorf("kmp() = %v, want %v", got, tt.want)
			}
		})
	}
}
