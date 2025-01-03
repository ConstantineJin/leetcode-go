package main

import "testing"

func Test_waysToSplitArray(t *testing.T) {
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
			args:    args{nums: []int{10, 4, -8, 7}},
			wantAns: 2,
		},
		{
			name:    "Example2",
			args:    args{nums: []int{2, 3, 1, 0}},
			wantAns: 2,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if gotAns := waysToSplitArray(tt.args.nums); gotAns != tt.wantAns {
				t.Errorf("waysToSplitArray() = %v, want %v", gotAns, tt.wantAns)
			}
		})
	}
}
