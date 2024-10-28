package main

import "testing"

func Test_longestSquareStreak(t *testing.T) {
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
			args:    args{nums: []int{4, 3, 6, 16, 8, 2}},
			wantAns: 3,
		},
		{
			name:    "Example2",
			args:    args{nums: []int{2, 3, 5, 6, 7}},
			wantAns: -1,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if gotAns := longestSquareStreak(tt.args.nums); gotAns != tt.wantAns {
				t.Errorf("longestSquareStreak() = %v, want %v", gotAns, tt.wantAns)
			}
		})
	}
}
