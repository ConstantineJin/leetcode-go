package main

import "testing"

func Test_countPairs(t *testing.T) {
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
			args:    args{nums: []int{1023, 2310, 2130, 213}},
			wantAns: 4,
		},
		{
			name:    "Example2",
			args:    args{nums: []int{1, 10, 100}},
			wantAns: 3,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if gotAns := countPairs(tt.args.nums); gotAns != tt.wantAns {
				t.Errorf("countPairs() = %v, want %v", gotAns, tt.wantAns)
			}
		})
	}
}
