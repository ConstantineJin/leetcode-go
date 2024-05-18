package main

import "testing"

func Test_maxDivScore(t *testing.T) {
	type args struct {
		nums     []int
		divisors []int
	}
	tests := []struct {
		name    string
		args    args
		wantAns int
	}{
		{
			name: "Example1",
			args: args{
				nums:     []int{4, 7, 9, 3, 9},
				divisors: []int{5, 2, 3},
			},
			wantAns: 3,
		},
		{
			name: "Example2",
			args: args{
				nums:     []int{20, 14, 21, 10},
				divisors: []int{5, 7, 5},
			},
			wantAns: 5,
		},
		{
			name: "Example3",
			args: args{
				nums:     []int{12},
				divisors: []int{10, 16},
			},
			wantAns: 10,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if gotAns := maxDivScore(tt.args.nums, tt.args.divisors); gotAns != tt.wantAns {
				t.Errorf("maxDivScore() = %v, want %v", gotAns, tt.wantAns)
			}
		})
	}
}
