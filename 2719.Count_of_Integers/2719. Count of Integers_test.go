package main

import "testing"

func Test_count(t *testing.T) {
	type args struct {
		num1   string
		num2   string
		minSum int
		maxSum int
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{
			name: "Example1",
			args: args{
				num1:   "1",
				num2:   "12",
				minSum: 1,
				maxSum: 8,
			},
			want: 11,
		},
		{
			name: "Example2",
			args: args{
				num1:   "1",
				num2:   "5",
				minSum: 1,
				maxSum: 5,
			},
			want: 5,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := count(tt.args.num1, tt.args.num2, tt.args.minSum, tt.args.maxSum); got != tt.want {
				t.Errorf("count() = %v, want %v", got, tt.want)
			}
		})
	}
}
