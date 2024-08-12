package main

import "testing"

func Test_countOfPairs(t *testing.T) {
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
			args:    args{nums: []int{2, 3, 2}},
			wantAns: 4,
		},
		{
			name:    "Example2",
			args:    args{nums: []int{5, 5, 5, 5}},
			wantAns: 126,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if gotAns := countOfPairs(tt.args.nums); gotAns != tt.wantAns {
				t.Errorf("countOfPairs() = %v, want %v", gotAns, tt.wantAns)
			}
		})
	}
}
