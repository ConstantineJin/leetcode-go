package main

import "testing"

func Test_hasTrailingZeros(t *testing.T) {
	type args struct {
		nums []int
	}
	tests := []struct {
		name    string
		args    args
		wantAns bool
	}{
		{
			name:    "Example1",
			args:    args{nums: []int{1, 2, 3, 4, 5}},
			wantAns: true,
		},
		{
			name:    "Example2",
			args:    args{nums: []int{2, 4, 88, 16}},
			wantAns: true,
		},
		{
			name:    "Example3",
			args:    args{nums: []int{1, 3, 5, 7, 9}},
			wantAns: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if gotAns := hasTrailingZeros(tt.args.nums); gotAns != tt.wantAns {
				t.Errorf("hasTrailingZeros() = %v, want %v", gotAns, tt.wantAns)
			}
		})
	}
}
