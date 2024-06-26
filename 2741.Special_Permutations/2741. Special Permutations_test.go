package main

import "testing"

func Test_specialPerm(t *testing.T) {
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
			args:    args{nums: []int{2, 3, 6}},
			wantAns: 2,
		},
		{
			name:    "Example2",
			args:    args{nums: []int{1, 4, 3}},
			wantAns: 2,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if gotAns := specialPerm(tt.args.nums); gotAns != tt.wantAns {
				t.Errorf("specialPerm() = %v, want %v", gotAns, tt.wantAns)
			}
		})
	}
}
