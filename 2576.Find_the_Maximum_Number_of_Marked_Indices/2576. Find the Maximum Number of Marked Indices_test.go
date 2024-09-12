package main

import "testing"

func Test_maxNumOfMarkedIndices(t *testing.T) {
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
			args:    args{nums: []int{3, 5, 2, 4}},
			wantAns: 2,
		},
		{
			name:    "Example2",
			args:    args{nums: []int{9, 2, 5, 4}},
			wantAns: 4,
		},
		{
			name:    "Example3",
			args:    args{nums: []int{7, 6, 8}},
			wantAns: 0,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if gotAns := maxNumOfMarkedIndices(tt.args.nums); gotAns != tt.wantAns {
				t.Errorf("maxNumOfMarkedIndices() = %v, want %v", gotAns, tt.wantAns)
			}
		})
	}
}
