package main

import "testing"

func Test_maxCount(t *testing.T) {
	type args struct {
		banned []int
		n      int
		maxSum int
	}
	tests := []struct {
		name    string
		args    args
		wantAns int
	}{
		{
			name: "Example1",
			args: args{
				banned: []int{1, 6, 5},
				n:      5,
				maxSum: 6,
			},
			wantAns: 2,
		},
		{
			name: "Example2",
			args: args{
				banned: []int{1, 2, 3, 4, 5, 6, 7},
				n:      8,
				maxSum: 1,
			},
			wantAns: 0,
		},
		{
			name: "Example3",
			args: args{
				banned: []int{11},
				n:      7,
				maxSum: 50,
			},
			wantAns: 7,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if gotAns := maxCount(tt.args.banned, tt.args.n, tt.args.maxSum); gotAns != tt.wantAns {
				t.Errorf("maxCount() = %v, want %v", gotAns, tt.wantAns)
			}
		})
	}
}
