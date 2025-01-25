package main

import "testing"

func Test_minimumMoney(t *testing.T) {
	type args struct {
		transactions [][]int
	}
	tests := []struct {
		name string
		args args
		want int64
	}{
		{
			name: "Example1",
			args: args{transactions: [][]int{{2, 1}, {5, 0}, {4, 2}}},
			want: 10,
		},
		{
			name: "Example2",
			args: args{transactions: [][]int{{3, 0}, {0, 3}}},
			want: 3,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := minimumMoney(tt.args.transactions); got != tt.want {
				t.Errorf("minimumMoney() = %v, want %v", got, tt.want)
			}
		})
	}
}
