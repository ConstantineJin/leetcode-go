package main

import "testing"

func Test_minChanges(t *testing.T) {
	type args struct {
		nums []int
		k    int
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{
			name: "Example1",
			args: args{
				nums: []int{1, 0, 1, 2, 4, 3},
				k:    4,
			},
			want: 2,
		},
		{
			name: "Example2",
			args: args{
				nums: []int{0, 1, 2, 3, 3, 6, 5, 4},
				k:    6,
			},
			want: 2,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := minChanges(tt.args.nums, tt.args.k); got != tt.want {
				t.Errorf("minChanges() = %v, want %v", got, tt.want)
			}
		})
	}
}
