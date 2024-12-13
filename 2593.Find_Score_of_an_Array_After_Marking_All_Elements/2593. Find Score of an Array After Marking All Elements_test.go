package main

import "testing"

func Test_findScore(t *testing.T) {
	type args struct {
		nums []int
	}
	tests := []struct {
		name    string
		args    args
		wantAns int64
	}{
		{
			name:    "Example1",
			args:    args{nums: []int{2, 1, 3, 4, 5, 2}},
			wantAns: 7,
		},
		{
			name:    "Example2",
			args:    args{nums: []int{2, 3, 5, 1, 3, 2}},
			wantAns: 5,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if gotAns := findScore(tt.args.nums); gotAns != tt.wantAns {
				t.Errorf("findScore() = %v, want %v", gotAns, tt.wantAns)
			}
		})
	}
}
