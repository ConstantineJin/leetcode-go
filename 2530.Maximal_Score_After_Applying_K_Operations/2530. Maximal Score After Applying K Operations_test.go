package main

import "testing"

func Test_maxKelements(t *testing.T) {
	type args struct {
		nums []int
		k    int
	}
	tests := []struct {
		name    string
		args    args
		wantSum int64
	}{
		{
			name: "Example1",
			args: args{
				nums: []int{10, 10, 10, 10, 10},
				k:    5,
			},
			wantSum: 50,
		},
		{
			name: "Example2",
			args: args{
				nums: []int{1, 10, 3, 3, 3},
				k:    3,
			},
			wantSum: 17,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if gotSum := maxKelements(tt.args.nums, tt.args.k); gotSum != tt.wantSum {
				t.Errorf("maxKelements() = %v, want %v", gotSum, tt.wantSum)
			}
		})
	}
}
