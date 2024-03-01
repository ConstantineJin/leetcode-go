package main

import "testing"

func Test_validPartition(t *testing.T) {
	type args struct {
		nums []int
	}
	tests := []struct {
		name string
		args args
		want bool
	}{
		{
			name: "Example1",
			args: args{nums: []int{4, 4, 4, 5, 6}},
			want: true,
		},
		{
			name: "Example2",
			args: args{nums: []int{1, 1, 1, 2}},
			want: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := validPartition(tt.args.nums); got != tt.want {
				t.Errorf("validPartition() = %v, want %v", got, tt.want)
			}
		})
	}
}
