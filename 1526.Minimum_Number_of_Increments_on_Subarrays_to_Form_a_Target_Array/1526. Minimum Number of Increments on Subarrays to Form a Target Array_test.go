package main

import "testing"

func Test_minNumberOperations(t *testing.T) {
	type args struct {
		target []int
	}
	tests := []struct {
		name    string
		args    args
		wantAns int
	}{
		{
			name:    "Example1",
			args:    args{target: []int{1, 2, 3, 2, 1}},
			wantAns: 3,
		},
		{
			name:    "Example2",
			args:    args{target: []int{3, 1, 1, 2}},
			wantAns: 4,
		},
		{
			name:    "Example3",
			args:    args{target: []int{3, 1, 5, 4, 2}},
			wantAns: 7,
		},
		{
			name:    "Example4",
			args:    args{target: []int{1, 1, 1, 1}},
			wantAns: 1,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if gotAns := minNumberOperations(tt.args.target); gotAns != tt.wantAns {
				t.Errorf("minNumberOperations() = %v, want %v", gotAns, tt.wantAns)
			}
		})
	}
}
