package main

import "testing"

func Test_maxGoodNumber(t *testing.T) {
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
			args:    args{nums: []int{1, 2, 3}},
			wantAns: 30,
		},
		{
			name:    "Example2",
			args:    args{nums: []int{2, 8, 16}},
			wantAns: 1296,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if gotAns := maxGoodNumber(tt.args.nums); gotAns != tt.wantAns {
				t.Errorf("maxGoodNumber() = %v, want %v", gotAns, tt.wantAns)
			}
		})
	}
}
