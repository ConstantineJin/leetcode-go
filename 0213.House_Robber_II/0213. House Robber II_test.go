package main

import "testing"

func Test_rob(t *testing.T) {
	type args struct {
		nums []int
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{
			name: "Example1",
			args: args{nums: []int{2, 3, 2}},
			want: 3,
		},
		{
			name: "Example2",
			args: args{nums: []int{1, 2, 3, 1}},
			want: 4,
		},
		{
			name: "Example3",
			args: args{nums: []int{1, 2, 3}},
			want: 3,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := rob(tt.args.nums); got != tt.want {
				t.Errorf("rob() = %v, want %v", got, tt.want)
			}
		})
	}
}
