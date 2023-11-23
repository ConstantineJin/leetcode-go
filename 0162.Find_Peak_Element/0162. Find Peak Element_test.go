package main

import "testing"

func Test_findPeakElement(t *testing.T) {
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
			args:    args{nums: []int{1, 2, 3, 1}},
			wantAns: 2,
		},
		{
			name:    "Example2",
			args:    args{nums: []int{1, 2, 1, 3, 5, 6, 4}},
			wantAns: 5,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if gotAns := findPeakElement(tt.args.nums); gotAns != tt.wantAns {
				t.Errorf("findPeakElement() = %v, want %v", gotAns, tt.wantAns)
			}
		})
	}
}
