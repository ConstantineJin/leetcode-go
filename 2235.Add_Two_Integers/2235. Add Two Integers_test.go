package main

import "testing"

func Test_sum(t *testing.T) {
	type args struct {
		num1 int
		num2 int
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{
			name: "Example1",
			args: args{
				num1: 12,
				num2: 5,
			},
			want: 17,
		},
		{
			name: "Example2",
			args: args{
				num1: -10,
				num2: 4,
			},
			want: -6,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := sum(tt.args.num1, tt.args.num2); got != tt.want {
				t.Errorf("sum() = %v, want %v", got, tt.want)
			}
		})
	}
}
