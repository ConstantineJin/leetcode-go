package main

import "testing"

func Test_differenceOfSum(t *testing.T) {
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
			args:    args{nums: []int{1, 15, 6, 3}},
			wantAns: 9,
		},
		{
			name:    "Example2",
			args:    args{nums: []int{1, 2, 3, 4}},
			wantAns: 0,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if gotAns := differenceOfSum(tt.args.nums); gotAns != tt.wantAns {
				t.Errorf("differenceOfSum() = %v, want %v", gotAns, tt.wantAns)
			}
		})
	}
}
