package main

import "testing"

func Test_maxWidthRamp(t *testing.T) {
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
			args:    args{nums: []int{6, 0, 8, 2, 1, 5}},
			wantAns: 4,
		},
		{
			name:    "Example2",
			args:    args{nums: []int{9, 8, 1, 0, 1, 9, 4, 0, 4, 1}},
			wantAns: 7,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if gotAns := maxWidthRamp(tt.args.nums); gotAns != tt.wantAns {
				t.Errorf("maxWidthRamp() = %v, want %v", gotAns, tt.wantAns)
			}
		})
	}
}
