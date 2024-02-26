package main

import "testing"

func Test_earliestSecondToMarkIndices(t *testing.T) {
	type args struct {
		nums          []int
		changeIndices []int
	}
	tests := []struct {
		name    string
		args    args
		wantAns int
	}{
		{
			name: "Example1",
			args: args{
				nums:          []int{2, 2, 0},
				changeIndices: []int{2, 2, 2, 2, 3, 2, 2, 1},
			},
			wantAns: 8,
		},
		{
			name: "Example2",
			args: args{
				nums:          []int{1, 3},
				changeIndices: []int{1, 1, 1, 2, 1, 1, 1},
			},
			wantAns: 6,
		},
		{
			name: "Example3",
			args: args{
				nums:          []int{0, 1},
				changeIndices: []int{2, 2, 2},
			},
			wantAns: -1,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if gotAns := earliestSecondToMarkIndices(tt.args.nums, tt.args.changeIndices); gotAns != tt.wantAns {
				t.Errorf("earliestSecondToMarkIndices() = %v, want %v", gotAns, tt.wantAns)
			}
		})
	}
}
