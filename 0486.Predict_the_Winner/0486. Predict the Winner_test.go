package main

import "testing"

func Test_predictTheWinner(t *testing.T) {
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
			args: args{nums: []int{1, 5, 2}},
			want: false,
		},
		{
			name: "Example2",
			args: args{nums: []int{1, 5, 233, 7}},
			want: true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := predictTheWinner(tt.args.nums); got != tt.want {
				t.Errorf("predictTheWinner() = %v, want %v", got, tt.want)
			}
		})
	}
}
