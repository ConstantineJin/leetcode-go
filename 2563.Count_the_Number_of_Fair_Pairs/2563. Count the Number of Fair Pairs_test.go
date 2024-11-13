package main

import "testing"

func Test_countFairPairs(t *testing.T) {
	type args struct {
		nums  []int
		lower int
		upper int
	}
	tests := []struct {
		name    string
		args    args
		wantAns int64
	}{
		{
			name: "Example1",
			args: args{
				nums:  []int{0, 1, 7, 4, 4, 5},
				lower: 3,
				upper: 6,
			},
			wantAns: 6,
		},
		{
			name: "Example2",
			args: args{
				nums:  []int{1, 7, 9, 2, 5},
				lower: 11,
				upper: 11,
			},
			wantAns: 1,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if gotAns := countFairPairs(tt.args.nums, tt.args.lower, tt.args.upper); gotAns != tt.wantAns {
				t.Errorf("countFairPairs() = %v, want %v", gotAns, tt.wantAns)
			}
		})
	}
}
