package main

import "testing"

func Test_maxCoins(t *testing.T) {
	type args struct {
		nums []int
	}
	tests := []struct {
		name    string
		args    args
		wantAns int
	}{
		{
			name:    "Example1",
			args:    args{nums: []int{3, 1, 5, 8}},
			wantAns: 167,
		},
		{
			name:    "Example2",
			args:    args{nums: []int{1, 5}},
			wantAns: 10,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if gotAns := maxCoins(tt.args.nums); gotAns != tt.wantAns {
				t.Errorf("maxCoins() = %v, want %v", gotAns, tt.wantAns)
			}
		})
	}
}
