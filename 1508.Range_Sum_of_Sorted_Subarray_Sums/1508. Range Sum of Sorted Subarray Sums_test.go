package main

import "testing"

func Test_rangeSum(t *testing.T) {
	type args struct {
		nums  []int
		n     int
		left  int
		right int
	}
	tests := []struct {
		name    string
		args    args
		wantAns int
	}{
		{
			name: "Example1",
			args: args{
				nums:  []int{1, 2, 3, 4},
				n:     4,
				left:  1,
				right: 5,
			},
			wantAns: 13,
		},
		{
			name: "Example2",
			args: args{
				nums:  []int{1, 2, 3, 4},
				n:     4,
				left:  3,
				right: 4,
			},
			wantAns: 6,
		},
		{
			name: "Example3",
			args: args{
				nums:  []int{1, 2, 3, 4},
				n:     4,
				left:  1,
				right: 10,
			},
			wantAns: 50,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if gotAns := rangeSum(tt.args.nums, tt.args.n, tt.args.left, tt.args.right); gotAns != tt.wantAns {
				t.Errorf("rangeSum() = %v, want %v", gotAns, tt.wantAns)
			}
		})
	}
}
