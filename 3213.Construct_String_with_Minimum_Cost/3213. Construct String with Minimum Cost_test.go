package main

import "testing"

func Test_minimumCost(t *testing.T) {
	type args struct {
		target string
		words  []string
		costs  []int
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{
			name: "Example1",
			args: args{
				target: "abcdef",
				words:  []string{"abdef", "abc", "d", "def", "ef"},
				costs:  []int{100, 1, 1, 10, 5},
			},
			want: 7,
		},
		{
			name: "Example2",
			args: args{
				target: "aaaa",
				words:  []string{"z", "zz", "zzz"},
				costs:  []int{1, 10, 100},
			},
			want: -1,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := minimumCost(tt.args.target, tt.args.words, tt.args.costs); got != tt.want {
				t.Errorf("minimumCost() = %v, want %v", got, tt.want)
			}
		})
	}
}
