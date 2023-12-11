package main

import "testing"

func Test_sumSubarrayMins(t *testing.T) {
	type args struct {
		arr []int
	}
	tests := []struct {
		name    string
		args    args
		wantAns int
	}{
		{
			name:    "Example1",
			args:    args{arr: []int{3, 1, 2, 4}},
			wantAns: 17,
		},
		{
			name:    "Example2",
			args:    args{arr: []int{11, 81, 94, 43, 3}},
			wantAns: 444,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if gotAns := sumSubarrayMins(tt.args.arr); gotAns != tt.wantAns {
				t.Errorf("sumSubarrayMins() = %v, want %v", gotAns, tt.wantAns)
			}
		})
	}
}
