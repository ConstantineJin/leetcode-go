package main

import "testing"

func Test_waysToSplit(t *testing.T) {
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
			args:    args{nums: []int{1, 1, 1}},
			wantAns: 1,
		},
		{
			name:    "Example2",
			args:    args{nums: []int{1, 2, 2, 2, 5, 0}},
			wantAns: 3,
		},
		{
			name:    "Example3",
			args:    args{nums: []int{3, 2, 1}},
			wantAns: 0,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if gotAns := waysToSplit(tt.args.nums); gotAns != tt.wantAns {
				t.Errorf("waysToSplit() = %v, want %v", gotAns, tt.wantAns)
			}
		})
	}
}
