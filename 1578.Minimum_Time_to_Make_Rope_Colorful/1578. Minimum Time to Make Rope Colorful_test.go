package main

import "testing"

func Test_minCost(t *testing.T) {
	type args struct {
		colors     string
		neededTime []int
	}
	tests := []struct {
		name    string
		args    args
		wantAns int
	}{
		{
			name: "Example1",
			args: args{
				colors:     "abaac",
				neededTime: []int{1, 2, 3, 4, 5},
			},
			wantAns: 3,
		},
		{
			name: "Example2",
			args: args{
				colors:     "abc",
				neededTime: []int{1, 2, 3},
			},
			wantAns: 0,
		},
		{
			name: "Example3",
			args: args{
				colors:     "aabaa",
				neededTime: []int{1, 2, 3, 4, 1},
			},
			wantAns: 2,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if gotAns := minCost(tt.args.colors, tt.args.neededTime); gotAns != tt.wantAns {
				t.Errorf("minCost() = %v, want %v", gotAns, tt.wantAns)
			}
		})
	}
}
