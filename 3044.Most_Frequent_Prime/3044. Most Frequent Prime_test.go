package main

import "testing"

func Test_mostFrequentPrime(t *testing.T) {
	type args struct {
		mat [][]int
	}
	tests := []struct {
		name    string
		args    args
		wantAns int
	}{
		{
			name:    "Example1",
			args:    args{mat: [][]int{{1, 1}, {9, 9}, {1, 1}}},
			wantAns: 19,
		},
		{
			name:    "Example2",
			args:    args{mat: [][]int{{7}}},
			wantAns: -1,
		},
		{
			name:    "Example3",
			args:    args{mat: [][]int{{9, 7, 8}, {4, 6, 5}, {2, 8, 6}}},
			wantAns: 97,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if gotAns := mostFrequentPrime(tt.args.mat); gotAns != tt.wantAns {
				t.Errorf("mostFrequentPrime() = %v, want %v", gotAns, tt.wantAns)
			}
		})
	}
}
